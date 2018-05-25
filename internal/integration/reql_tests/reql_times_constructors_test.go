// Code generated by gen_tests.py and process_polyglot.py.
// Do not edit this file directly.
// The template for this file is located at:
// ../template.go.tpl
package reql_tests

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	r "gopkg.in/rebirthdb/rebirthdb-go.v4"
	"gopkg.in/rebirthdb/rebirthdb-go.v4/internal/compare"
)

// Test basic time arithmetic
func TestTimesConstructorsSuite(t *testing.T) {
	suite.Run(t, new(TimesConstructorsSuite))
}

type TimesConstructorsSuite struct {
	suite.Suite

	session *r.Session
}

func (suite *TimesConstructorsSuite) SetupTest() {
	suite.T().Log("Setting up TimesConstructorsSuite")
	// Use imports to prevent errors
	_ = time.Time{}
	_ = compare.AnythingIsFine

	session, err := r.Connect(r.ConnectOpts{
		Address: url,
	})
	suite.Require().NoError(err, "Error returned when connecting to server")
	suite.session = session

	r.DBDrop("test").Exec(suite.session)
	err = r.DBCreate("test").Exec(suite.session)
	suite.Require().NoError(err)
	err = r.DB("test").Wait().Exec(suite.session)
	suite.Require().NoError(err)

}

func (suite *TimesConstructorsSuite) TearDownSuite() {
	suite.T().Log("Tearing down TimesConstructorsSuite")

	if suite.session != nil {
		r.DB("rethinkdb").Table("_debug_scratch").Delete().Exec(suite.session)
		r.DBDrop("test").Exec(suite.session)

		suite.session.Close()
	}
}

func (suite *TimesConstructorsSuite) TestCases() {
	suite.T().Log("Running TimesConstructorsSuite: Test basic time arithmetic")

	{
		// times/constructors.yaml line #5
		/* datetime.fromtimestamp(896571240, r.ast.RqlTzinfo('00:00')) */
		var expected_ time.Time = Ast.Fromtimestamp(896571240, Ast.RqlTzinfo("00:00"))
		/* r.expr(r.epoch_time(896571240)) */

		suite.T().Log("About to run line #5: r.Expr(r.EpochTime(896571240))")

		runAndAssert(suite.Suite, expected_, r.Expr(r.EpochTime(896571240)), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #5")
	}

	{
		// times/constructors.yaml line #11
		/* {'stuff':datetime.fromtimestamp(896571240, r.ast.RqlTzinfo('00:00')), 'more':[datetime.fromtimestamp(996571240, r.ast.RqlTzinfo('00:00'))]} */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"stuff": Ast.Fromtimestamp(896571240, Ast.RqlTzinfo("00:00")), "more": []interface{}{Ast.Fromtimestamp(996571240, Ast.RqlTzinfo("00:00"))}}
		/* r.expr({'stuff':r.epoch_time(896571240), 'more':[r.epoch_time(996571240)]}) */

		suite.T().Log("About to run line #11: r.Expr(map[interface{}]interface{}{'stuff': r.EpochTime(896571240), 'more': []interface{}{r.EpochTime(996571240)}, })")

		runAndAssert(suite.Suite, expected_, r.Expr(map[interface{}]interface{}{"stuff": r.EpochTime(896571240), "more": []interface{}{r.EpochTime(996571240)}}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #11")
	}

	{
		// times/constructors.yaml line #17
		/* [datetime.fromtimestamp(796571240, r.ast.RqlTzinfo('00:00')), datetime.fromtimestamp(896571240, r.ast.RqlTzinfo('00:00')), {'stuff':datetime.fromtimestamp(996571240, r.ast.RqlTzinfo('00:00'))}] */
		var expected_ []interface{} = []interface{}{Ast.Fromtimestamp(796571240, Ast.RqlTzinfo("00:00")), Ast.Fromtimestamp(896571240, Ast.RqlTzinfo("00:00")), map[interface{}]interface{}{"stuff": Ast.Fromtimestamp(996571240, Ast.RqlTzinfo("00:00"))}}
		/* r.expr([r.epoch_time(796571240), r.epoch_time(896571240), {'stuff':r.epoch_time(996571240)}]) */

		suite.T().Log("About to run line #17: r.Expr([]interface{}{r.EpochTime(796571240), r.EpochTime(896571240), map[interface{}]interface{}{'stuff': r.EpochTime(996571240), }})")

		runAndAssert(suite.Suite, expected_, r.Expr([]interface{}{r.EpochTime(796571240), r.EpochTime(896571240), map[interface{}]interface{}{"stuff": r.EpochTime(996571240)}}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #17")
	}

	{
		// times/constructors.yaml line #23
		/* {'nested':{'time':datetime.fromtimestamp(896571240, r.ast.RqlTzinfo('00:00'))}} */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"nested": map[interface{}]interface{}{"time": Ast.Fromtimestamp(896571240, Ast.RqlTzinfo("00:00"))}}
		/* r.expr({'nested':{'time':r.epoch_time(896571240)}}) */

		suite.T().Log("About to run line #23: r.Expr(map[interface{}]interface{}{'nested': map[interface{}]interface{}{'time': r.EpochTime(896571240), }, })")

		runAndAssert(suite.Suite, expected_, r.Expr(map[interface{}]interface{}{"nested": map[interface{}]interface{}{"time": r.EpochTime(896571240)}}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #23")
	}

	{
		// times/constructors.yaml line #29
		/* [1, "two", ["a", datetime.fromtimestamp(896571240, r.ast.RqlTzinfo('00:00')), 3]] */
		var expected_ []interface{} = []interface{}{1, "two", []interface{}{"a", Ast.Fromtimestamp(896571240, Ast.RqlTzinfo("00:00")), 3}}
		/* r.expr([1, "two", ["a", r.epoch_time(896571240), 3]]) */

		suite.T().Log("About to run line #29: r.Expr([]interface{}{1, 'two', []interface{}{'a', r.EpochTime(896571240), 3}})")

		runAndAssert(suite.Suite, expected_, r.Expr([]interface{}{1, "two", []interface{}{"a", r.EpochTime(896571240), 3}}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #29")
	}

	{
		// times/constructors.yaml line #35
		/* 1 */
		var expected_ int = 1
		/* r.epoch_time(1).to_epoch_time() */

		suite.T().Log("About to run line #35: r.EpochTime(1).ToEpochTime()")

		runAndAssert(suite.Suite, expected_, r.EpochTime(1).ToEpochTime(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #35")
	}

	{
		// times/constructors.yaml line #37
		/* -1 */
		var expected_ int = -1
		/* r.epoch_time(-1).to_epoch_time() */

		suite.T().Log("About to run line #37: r.EpochTime(-1).ToEpochTime()")

		runAndAssert(suite.Suite, expected_, r.EpochTime(-1).ToEpochTime(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #37")
	}

	{
		// times/constructors.yaml line #39
		/* 1.444 */
		var expected_ float64 = 1.444
		/* r.epoch_time(1.4444445).to_epoch_time() */

		suite.T().Log("About to run line #39: r.EpochTime(1.4444445).ToEpochTime()")

		runAndAssert(suite.Suite, expected_, r.EpochTime(1.4444445).ToEpochTime(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #39")
	}

	{
		// times/constructors.yaml line #42
		/* "1970-01-01T00:00:01.444+00:00" */
		var expected_ string = "1970-01-01T00:00:01.444+00:00"
		/* r.epoch_time(1.4444445).to_iso8601() */

		suite.T().Log("About to run line #42: r.EpochTime(1.4444445).ToISO8601()")

		runAndAssert(suite.Suite, expected_, r.EpochTime(1.4444445).ToISO8601(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #42")
	}

	{
		// times/constructors.yaml line #45
		/* 1.444 */
		var expected_ float64 = 1.444
		/* r.epoch_time(1.4444445).seconds() */

		suite.T().Log("About to run line #45: r.EpochTime(1.4444445).Seconds()")

		runAndAssert(suite.Suite, expected_, r.EpochTime(1.4444445).Seconds(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #45")
	}

	{
		// times/constructors.yaml line #48
		/* 10000 */
		var expected_ int = 10000
		/* r.epoch_time(253430000000).year() */

		suite.T().Log("About to run line #48: r.EpochTime(253430000000).Year()")

		runAndAssert(suite.Suite, expected_, r.EpochTime(253430000000).Year(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #48")
	}

	{
		// times/constructors.yaml line #50
		/* err("ReqlQueryLogicError", "Year `10000` out of valid ISO 8601 range [0, 9999].", []) */
		var expected_ Err = err("ReqlQueryLogicError", "Year `10000` out of valid ISO 8601 range [0, 9999].")
		/* r.epoch_time(253430000000).to_iso8601() */

		suite.T().Log("About to run line #50: r.EpochTime(253430000000).ToISO8601()")

		runAndAssert(suite.Suite, expected_, r.EpochTime(253430000000).ToISO8601(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #50")
	}

	{
		// times/constructors.yaml line #53
		/* err("ReqlQueryLogicError", "Error in time logic: Year is out of valid range: 1400..10000.", []) */
		var expected_ Err = err("ReqlQueryLogicError", "Error in time logic: Year is out of valid range: 1400..10000.")
		/* r.epoch_time(253440000000).year() */

		suite.T().Log("About to run line #53: r.EpochTime(253440000000).Year()")

		runAndAssert(suite.Suite, expected_, r.EpochTime(253440000000).Year(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #53")
	}

	{
		// times/constructors.yaml line #55
		/* 253440000000 */
		var expected_ int = 253440000000
		/* r.epoch_time(253440000000).to_epoch_time() */

		suite.T().Log("About to run line #55: r.EpochTime(253440000000).ToEpochTime()")

		runAndAssert(suite.Suite, expected_, r.EpochTime(253440000000).ToEpochTime(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #55")
	}

	{
		// times/constructors.yaml line #57
		/* 1400 */
		var expected_ int = 1400
		/* r.epoch_time(-17980000000).year() */

		suite.T().Log("About to run line #57: r.EpochTime(-17980000000).Year()")

		runAndAssert(suite.Suite, expected_, r.EpochTime(-17980000000).Year(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #57")
	}

	{
		// times/constructors.yaml line #59
		/* err("ReqlQueryLogicError", "Error in time logic: Year is out of valid range: 1400..10000.", []) */
		var expected_ Err = err("ReqlQueryLogicError", "Error in time logic: Year is out of valid range: 1400..10000.")
		/* r.epoch_time(-17990000000).year() */

		suite.T().Log("About to run line #59: r.EpochTime(-17990000000).Year()")

		runAndAssert(suite.Suite, expected_, r.EpochTime(-17990000000).Year(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #59")
	}

	{
		// times/constructors.yaml line #61
		/* -17990000000 */
		var expected_ int = -17990000000
		/* r.epoch_time(-17990000000).to_epoch_time() */

		suite.T().Log("About to run line #61: r.EpochTime(-17990000000).ToEpochTime()")

		runAndAssert(suite.Suite, expected_, r.EpochTime(-17990000000).ToEpochTime(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #61")
	}

	// times/constructors.yaml line #65
	// cdate = "2013-01-01"
	suite.T().Log("Possibly executing: var cdate string = '2013-01-01'")

	cdate := "2013-01-01"
	_ = cdate // Prevent any noused variable errors

	// times/constructors.yaml line #66
	// dates = ["2013", "2013-01", "2013-01-01", "20130101", "2013-001", "2013001"]
	suite.T().Log("Possibly executing: var dates []interface{} = []interface{}{'2013', '2013-01', '2013-01-01', '20130101', '2013-001', '2013001'}")

	dates := []interface{}{"2013", "2013-01", "2013-01-01", "20130101", "2013-001", "2013001"}
	_ = dates // Prevent any noused variable errors

	// times/constructors.yaml line #67
	// ctime = "13:00:00"
	suite.T().Log("Possibly executing: var ctime string = '13:00:00'")

	ctime := "13:00:00"
	_ = ctime // Prevent any noused variable errors

	// times/constructors.yaml line #68
	// times = ["13", "13:00", "1300", "13:00:00", "13:00:00.000000", "130000.000000"]
	suite.T().Log("Possibly executing: var times []interface{} = []interface{}{'13', '13:00', '1300', '13:00:00', '13:00:00.000000', '130000.000000'}")

	times := []interface{}{"13", "13:00", "1300", "13:00:00", "13:00:00.000000", "130000.000000"}
	_ = times // Prevent any noused variable errors

	// times/constructors.yaml line #69
	// ctz = "+00:00"
	suite.T().Log("Possibly executing: var ctz string = '+00:00'")

	ctz := "+00:00"
	_ = ctz // Prevent any noused variable errors

	// times/constructors.yaml line #70
	// tzs = ["Z", "+00", "+0000", "+00:00"]
	suite.T().Log("Possibly executing: var tzs []interface{} = []interface{}{'Z', '+00', '+0000', '+00:00'}")

	tzs := []interface{}{"Z", "+00", "+0000", "+00:00"}
	_ = tzs // Prevent any noused variable errors

	// times/constructors.yaml line #71
	// cdt = [cdate+"T"+ctime+ctz]
	suite.T().Log("Possibly executing: var cdt []interface{} = []interface{}{cdate + 'T' + ctime + ctz}")

	cdt := []interface{}{cdate + "T" + ctime + ctz}
	_ = cdt // Prevent any noused variable errors

	// times/constructors.yaml line #81
	// bad_dates = ["201301", "2013-0101", "2a13", "2013+01", "2013-01-01.1"]
	suite.T().Log("Possibly executing: var bad_dates []interface{} = []interface{}{'201301', '2013-0101', '2a13', '2013+01', '2013-01-01.1'}")

	bad_dates := []interface{}{"201301", "2013-0101", "2a13", "2013+01", "2013-01-01.1"}
	_ = bad_dates // Prevent any noused variable errors

	// times/constructors.yaml line #82
	// bad_times = ["a3", "13:0000", "13:000", "13:00.00", "130000.00000000a"]
	suite.T().Log("Possibly executing: var bad_times []interface{} = []interface{}{'a3', '13:0000', '13:000', '13:00.00', '130000.00000000a'}")

	bad_times := []interface{}{"a3", "13:0000", "13:000", "13:00.00", "130000.00000000a"}
	_ = bad_times // Prevent any noused variable errors

	// times/constructors.yaml line #83
	// bad_tzs = ["X", "-7", "-07:-1", "+07+01", "PST", "UTC", "Z+00"]
	suite.T().Log("Possibly executing: var bad_tzs []interface{} = []interface{}{'X', '-7', '-07:-1', '+07+01', 'PST', 'UTC', 'Z+00'}")

	bad_tzs := []interface{}{"X", "-7", "-07:-1", "+07+01", "PST", "UTC", "Z+00"}
	_ = bad_tzs // Prevent any noused variable errors

}
