package userpasswd

import (
	"github.com/khulnasoft/khulnasoft/internal/txemail"
	"github.com/khulnasoft/khulnasoft/internal/txemail/txtypes"
)

type SetPasswordEmailTemplateData struct {
	Username string
	URL      string
	Host     string
}

var defaultSetPasswordEmailTemplate = txemail.MustValidate(txtypes.Templates{
	Subject: `Set your Khulnasoft password ({{.Host}})`,
	Text: `
Your administrator created an account for you on Khulnasoft ({{.Host}}).

To set the password for {{.Username}} on Khulnasoft, follow this link:

  {{.URL}}
`,
	HTML: `
<p>
  Your administrator created an account for you on Khulnasoft ({{.Host}}).
</p>

<p><strong><a href="{{.URL}}">Set password for {{.Username}}</a></strong></p>
`,
})

var defaultResetPasswordEmailTemplates = txemail.MustValidate(txtypes.Templates{
	Subject: `Reset your Khulnasoft password ({{.Host}})`,
	Text: `
Somebody (likely you) requested a password reset for the user {{.Username}} on Khulnasoft ({{.Host}}).

To reset the password for {{.Username}} on Khulnasoft, follow this link:

  {{.URL}}
`,
	HTML: `
<p>
  Somebody (likely you) requested a password reset for <strong>{{.Username}}</strong>
  on Khulnasoft ({{.Host}}).
</p>

<p><strong><a href="{{.URL}}">Reset password for {{.Username}}</a></strong></p>
`,
})
