package auth

import (
	"github.com/thomiceli/opengist/internal/web/context"
	webhandlers "github.com/thomiceli/opengist/internal/web/handlers"
)

const nextSessionKey = "next"

func nextFromRequest(ctx *context.Context) string {
	return webhandlers.SafeNext(ctx.FormValue(webhandlers.NextParam))
}

func nextFromSession(ctx *context.Context) string {
	target, _ := ctx.GetSession().Values[nextSessionKey].(string)
	return webhandlers.SafeNext(target)
}

func saveNext(ctx *context.Context, next string) {
	sess := ctx.GetSession()
	if next == "/" {
		delete(sess.Values, nextSessionKey)
	} else {
		sess.Values[nextSessionKey] = next
	}
	ctx.SaveSession(sess)
}

func deleteNext(ctx *context.Context) {
	sess := ctx.GetSession()
	delete(sess.Values, nextSessionKey)
	ctx.SaveSession(sess)
}

func loginRedirect(next string) string {
	return webhandlers.LoginRedirect(next)
}
