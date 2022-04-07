package banktest

import (
	"std"
)

type activity struct {
	caller   std.Address
	sent     std.Coins
	returned std.Coins
	time     std.Time
}

func (act *activity) String() string {
	return act.caller.String() + " " +
		act.sent.String() + " sent, " +
		act.returned.String() + " returned, at " +
		std.FormatTimestamp(act.time, "2006-01-02 3:04pm MST")
}

var latest [10]*activity

// Deposit will take the coins (to the realm's pkgaddr) or return them to user.
func Deposit(returnDenom string, returnAmount int64) string {
	std.AssertOriginCall()
	caller := std.GetOrigCaller()
	send := std.Coins{{returnDenom, returnAmount}}
	// record activity
	act := &activity{
		caller:   caller,
		sent:     std.GetOrigSend(),
		returned: send,
		time:     std.GetTimestamp(),
	}
	for i := len(latest) - 1; i >= 0; i-- {
		latest[i+1] = latest[i] // shift by +1.
	}
	latest[0] = act
	// return if any.
	if returnAmount > 0 {
		banker := std.GetBanker(std.BankerTypeOrigSend)
		pkgaddr := std.GetOrigPkgAddr()
		// TODO: use std.Coins constructors, this isn't generally safe.
		banker.SendCoins(pkgaddr, caller, send)
		return "returned!"
	} else {
		return "thank you!"
	}
}

func Render(path string) string {
	// get realm coins.
	banker := std.GetBanker(std.BankerTypeReadonly)
	coins := banker.GetCoins(std.GetOrigPkgAddr())

	// render
	res := ""
	res += "## recent activity\n"
	res += "\n"
	for _, act := range latest {
		if act == nil {
			break
		}
		res += " * " + act.String() + "\n"
	}
	res += "\n"
	res += "## total deposits\n"
	res += coins.String()
	return res
}
