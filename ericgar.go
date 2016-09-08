package ericgar

/*
import (
	"crypto"
	"crypto/ecdsa"
	"net/http"

	"github.com/xenolf/lego/acme"
)

// const letsEncryptURL = "https://acme-v01.api.letsencrypt.org/directory"
const letsEncryptURL = "https://acme-staging.api.letsencrypt.org/directory"

func init() {
	http.HandleFunc("/admin/newcert", newCert)
}

type acmeUser struct {
	Email string
	Reg   *acme.RegistrationResource
	key   *ecdsa.PrivateKey
}

func (au *acmeUser) GetEmail() string                            { return au.Email }
func (au *acmeUser) GetRegistration() *acme.RegistrationResource { return au.Reg }
func (au *acmeUser) GetPrivateKey() crypto.PrivateKey            { return au.key }

/*
func newCert(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	u := user.Current(ctx)
	if u == nil {
		url, _ := user.LoginURL(ctx, "/")
		fmt.Fprintf(w, `<a href="%s">Sign in or register</a>`, url)
		return
	}
	acme.HTTPClient = urlfetch.Client(ctx)
	au := &acmeUser{
		Email: u.
	}
	a, err := acme.NewClient(letsEncryptURL, au, acme.EC256)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	reg, err := a.Register()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := a.AgreeToTOS(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	certs, errMap := a.ObtainCertificate([]string{"ssh.ericgar.com"}, true, nil)
	if len(errMap) > 0 {
		http.Error(w, fmt.Sprintf("%+v", errMap), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "Cert:\n\n%v\n\n", certs.Certificate)
	fmt.Fprintf(w, "Private Key:\n\n%v", certs.PrivateKey)
}
*/
