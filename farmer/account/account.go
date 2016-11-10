package account

import (
	"crypto/x509"
	"fmt"
	"time"

	"github.com/conseweb/common/hdwallet"
	pb "github.com/conseweb/common/protos"
	"github.com/hyperledger/fabric/core/crypto/primitives"
	"github.com/op/go-logging"
	"github.com/spf13/viper"
	// "golang.org/x/crypto/bcrypt"
	"golang.org/x/net/context"
	// "google.golang.org/grpc"
)

const (
	defaultTimeout = time.Second * 3
)

var LanguageSupport = map[string]pb.PassphraseLanguage{
	"English":  pb.PassphraseLanguage_English,
	"简体中文":     pb.PassphraseLanguage_SimplifiedChinese,
	"繁體中文":     pb.PassphraseLanguage_TraditionalChinese,
	"日本語":      pb.PassphraseLanguage_JAPANESE,
	"español":  pb.PassphraseLanguage_SPANISH,
	"français": pb.PassphraseLanguage_FRENCH,
	"ITALIAN":  pb.PassphraseLanguage_ITALIAN,
}

type Account struct {
	ID       string                `json:"id"`
	NickName string                `json:"nicename"`
	Phone    string                `json:"phone"`
	Email    string                `json:"email"`
	Password string                `json:"password"`
	Lang     pb.PassphraseLanguage `json:"lang"`

	Passphrase string `json:"passphrase"`
	Wallet     *hdwallet.HDWallet

	Devices []Device

	logger *logging.Logger
}

func NewAccount(nickname, phone, email, pass, lang string) *Account {
	language := pb.PassphraseLanguage_English
	if l, ok := LanguageSupport[lang]; ok {
		language = l
	}

	dev_mode := viper.GetBool("daemon.dev")
	ph, hd := hdwallet.NewHDWallet(pass, language, dev_mode)
	return &Account{
		Phone:    phone,
		Email:    email,
		NickName: nickname,
		Password: pass,
		Lang:     language,

		Passphrase: ph,
		Wallet:     hd,
		logger:     logging.MustGetLogger("farmer"),
	}
}

func (a *Account) Registry(idpCli pb.IDPPClient) error {
	priv, err := primitives.NewECDSAKey()
	if err != nil {
		return err
	}

	pubraw, err := x509.MarshalPKIXPublicKey(&priv.PublicKey)
	if err != nil {
		return err
	}

	regUser := &pb.RegisterUserReq{
		SignUpType: pb.SignUpType_MOBILE,
		SignUp:     a.Phone,
		Nick:       a.NickName,
		Pass:       a.Password,
		UserType:   pb.UserType_NORMAL,

		Wpub: []byte(a.Wallet.Pub().String()),
		Spub: pubraw,
		Sign: []byte("ffff"),
	}
	if ok := checkPhone(regUser.SignUp); !ok {
		if ok := checkEmail(a.Email); !ok {
			return fmt.Errorf("Email and phone need at least one.")
		}
		regUser.SignUpType = pb.SignUpType_EMAIL
		regUser.SignUp = a.Email
	}

	resp, err := idpCli.RegisterUser(context.Background(), regUser)
	if err != nil {
		return err
	}
	if resp.GetError() != nil && resp.GetError().ErrorType != pb.ErrorType_NONE_ERROR {
		return resp.GetError()
	}

	ru := resp.GetUser()
	if ru == nil {
		return fmt.Errorf("got nil user.")
	}
	a.ID = ru.UserID

	a.Save()
	return nil
}

func Login(idpCli pb.IDPPClient, typ pb.SignInType, signup, password string) (a *Account, err error) {
	req := &pb.LoginUserReq{
		SignInType: typ,
		SignIn:     signup,
		Password:   password,
		Sign:       []byte("ffff"),
	}

	resp, err := idpCli.LoginUser(context.Background(), req)
	if err != nil {
		return nil, err
	}
	if resp.GetError() != nil && resp.GetError().ErrorType != pb.ErrorType_NONE_ERROR {
		return nil, resp.GetError()
	}

	ru := resp.GetUser()
	if ru == nil {
		return nil, fmt.Errorf("got nil user.")
	}
	a = &Account{
		ID:       ru.UserID,
		Phone:    ru.Mobile,
		Email:    ru.Email,
		NickName: ru.Nick,
		Devices:  []Device{},
	}
	for _, device := range ru.Devices {
		a.Devices = append(a.Devices, Device{Device: device})
	}
	a.Save()
	return a, nil
}

func (a *Account) BindDevice(idpCli pb.IDPPClient) error {
	priv, err := primitives.NewECDSAKey()
	if err != nil {
		return err
	}

	pubraw, err := x509.MarshalPKIXPublicKey(&priv.PublicKey)
	if err != nil {
		return err
	}

	wlt, err := a.Wallet.Child(0)
	if err != nil {
		return err
	}

	devReq := &pb.BindDeviceReq{
		UserID: a.ID,
		Os:     "linux",
		For:    pb.DeviceFor_FARMER,
		Mac:    "mac",
		Alias:  "haha",

		Wpub: []byte(wlt.Pub().String()),
		// device signature public key
		Spub: pubraw,
		// request signature, though has spub, but using user's spub signature this message, dont forget!
		Sign: []byte("ffff"),
	}
	_ = devReq
	return nil
}

func (a *Account) Logout() error {
	return nil
}

func (a *Account) Save() error {
	return nil
}

func (a *Account) Online(client pb.FarmerPublicClient) error {
	if a.ID == "" {
		return fmt.Errorf("account id required")
	}

	onlineReq := &pb.FarmerOnLineReq{FarmerID: a.ID}
	a.logger.Debugf("login with %+v", a)

	onlineRes, err := client.FarmerOnLine(context.Background(), onlineReq)
	if err != nil {
		return err
	}
	if onlineRes.Error != nil {
		a.logger.Errorf("login error: %#v", onlineRes.Error)
		return onlineRes.Error
	}

	return nil
}

func (a *Account) getSignInType() (st pb.SignInType, sv string) {
	st, sv = pb.SignInType_SI_MOBILE, a.Phone
	if a.ID != "" {
		st, sv = pb.SignInType_SI_USERID, a.ID
	} else if a.Email != "" {
		st, sv = pb.SignInType_SI_EMAIL, a.Email
	}
	return
}
