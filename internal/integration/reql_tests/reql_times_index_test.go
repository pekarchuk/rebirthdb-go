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

// secondary indexes on times
func TestTimesIndexSuite(t *testing.T) {
	suite.Run(t, new(TimesIndexSuite))
}

type TimesIndexSuite struct {
	suite.Suite

	session *r.Session
}

func (suite *TimesIndexSuite) SetupTest() {
	suite.T().Log("Setting up TimesIndexSuite")
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

	r.DB("test").TableDrop("tbl").Exec(suite.session)
	err = r.DB("test").TableCreate("tbl").Exec(suite.session)
	suite.Require().NoError(err)
	err = r.DB("test").Table("tbl").Wait().Exec(suite.session)
	suite.Require().NoError(err)
}

func (suite *TimesIndexSuite) TearDownSuite() {
	suite.T().Log("Tearing down TimesIndexSuite")

	if suite.session != nil {
		r.DB("rethinkdb").Table("_debug_scratch").Delete().Exec(suite.session)
		r.DB("test").TableDrop("tbl").Exec(suite.session)
		r.DBDrop("test").Exec(suite.session)

		suite.session.Close()
	}
}

func (suite *TimesIndexSuite) TestCases() {
	suite.T().Log("Running TimesIndexSuite: secondary indexes on times")

	tbl := r.DB("test").Table("tbl")
	_ = tbl // Prevent any noused variable errors

	// times/index.yaml line #7
	// ts={"timezone":"-07:00","epoch_time":1375445162.0872,"$reql_type$":"TIME"}
	suite.T().Log("Possibly executing: var ts map[interface{}]interface{} = map[interface{}]interface{}{'timezone': '-07:00', 'epoch_time': 1375445162.0872, '$reql_type$': 'TIME', }")

	ts := map[interface{}]interface{}{"timezone": "-07:00", "epoch_time": 1375445162.0872, "$reql_type$": "TIME"}
	_ = ts // Prevent any noused variable errors

	// times/index.yaml line #11
	// t1={"timezone":"-07:00","epoch_time":1375445163.0872,"$reql_type$":"TIME"}
	suite.T().Log("Possibly executing: var t1 map[interface{}]interface{} = map[interface{}]interface{}{'timezone': '-07:00', 'epoch_time': 1375445163.0872, '$reql_type$': 'TIME', }")

	t1 := map[interface{}]interface{}{"timezone": "-07:00", "epoch_time": 1375445163.0872, "$reql_type$": "TIME"}
	_ = t1 // Prevent any noused variable errors

	// times/index.yaml line #15
	// t2={"timezone":"-07:00","epoch_time":1375445163.08832,"$reql_type$":"TIME"}
	suite.T().Log("Possibly executing: var t2 map[interface{}]interface{} = map[interface{}]interface{}{'timezone': '-07:00', 'epoch_time': 1375445163.08832, '$reql_type$': 'TIME', }")

	t2 := map[interface{}]interface{}{"timezone": "-07:00", "epoch_time": 1375445163.08832, "$reql_type$": "TIME"}
	_ = t2 // Prevent any noused variable errors

	// times/index.yaml line #19
	// t3={"timezone":"-07:00","epoch_time":1375445163.08943,"$reql_type$":"TIME"}
	suite.T().Log("Possibly executing: var t3 map[interface{}]interface{} = map[interface{}]interface{}{'timezone': '-07:00', 'epoch_time': 1375445163.08943, '$reql_type$': 'TIME', }")

	t3 := map[interface{}]interface{}{"timezone": "-07:00", "epoch_time": 1375445163.08943, "$reql_type$": "TIME"}
	_ = t3 // Prevent any noused variable errors

	// times/index.yaml line #23
	// t4={"timezone":"-07:00","epoch_time":1375445163.09055,"$reql_type$":"TIME"}
	suite.T().Log("Possibly executing: var t4 map[interface{}]interface{} = map[interface{}]interface{}{'timezone': '-07:00', 'epoch_time': 1375445163.09055, '$reql_type$': 'TIME', }")

	t4 := map[interface{}]interface{}{"timezone": "-07:00", "epoch_time": 1375445163.09055, "$reql_type$": "TIME"}
	_ = t4 // Prevent any noused variable errors

	// times/index.yaml line #27
	// t5={"timezone":"-07:00","epoch_time":1375445163.09166,"$reql_type$":"TIME"}
	suite.T().Log("Possibly executing: var t5 map[interface{}]interface{} = map[interface{}]interface{}{'timezone': '-07:00', 'epoch_time': 1375445163.09166, '$reql_type$': 'TIME', }")

	t5 := map[interface{}]interface{}{"timezone": "-07:00", "epoch_time": 1375445163.09166, "$reql_type$": "TIME"}
	_ = t5 // Prevent any noused variable errors

	// times/index.yaml line #31
	// te={"timezone":"-07:00","epoch_time":1375445164.0872,"$reql_type$":"TIME"}
	suite.T().Log("Possibly executing: var te map[interface{}]interface{} = map[interface{}]interface{}{'timezone': '-07:00', 'epoch_time': 1375445164.0872, '$reql_type$': 'TIME', }")

	te := map[interface{}]interface{}{"timezone": "-07:00", "epoch_time": 1375445164.0872, "$reql_type$": "TIME"}
	_ = te // Prevent any noused variable errors

	// times/index.yaml line #36
	// trows = [{'id':t1}, {'id':t2}, {'id':t3}, {'id':t4}, {'id':t5}]
	suite.T().Log("Possibly executing: var trows []interface{} = []interface{}{map[interface{}]interface{}{'id': t1, }, map[interface{}]interface{}{'id': t2, }, map[interface{}]interface{}{'id': t3, }, map[interface{}]interface{}{'id': t4, }, map[interface{}]interface{}{'id': t5, }}")

	trows := []interface{}{map[interface{}]interface{}{"id": t1}, map[interface{}]interface{}{"id": t2}, map[interface{}]interface{}{"id": t3}, map[interface{}]interface{}{"id": t4}, map[interface{}]interface{}{"id": t5}}
	_ = trows // Prevent any noused variable errors

	{
		// times/index.yaml line #37
		/* 5 */
		var expected_ int = 5
		/* tbl.insert(trows)['inserted'] */

		suite.T().Log("About to run line #37: tbl.Insert(trows).AtIndex('inserted')")

		runAndAssert(suite.Suite, expected_, tbl.Insert(trows).AtIndex("inserted"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #37")
	}

	// times/index.yaml line #41
	// bad_insert = [{'id':r.expr(t1).in_timezone("Z")}]
	suite.T().Log("Possibly executing: var bad_insert []interface{} = []interface{}{map[interface{}]interface{}{'id': r.Expr(t1).InTimezone('Z'), }}")

	bad_insert := []interface{}{map[interface{}]interface{}{"id": r.Expr(t1).InTimezone("Z")}}
	_ = bad_insert // Prevent any noused variable errors

	{
		// times/index.yaml line #42
		/* ("Duplicate primary key `id`:\n{\n\t\"id\":\t{\n\t\t\"$reql_type$\":\t\"TIME\",\n\t\t\"epoch_time\":\t1375445163.087,\n\t\t\"timezone\":\t\"-07:00\"\n\t}\n}\n{\n\t\"id\":\t{\n\t\t\"$reql_type$\":\t\"TIME\",\n\t\t\"epoch_time\":\t1375445163.087,\n\t\t\"timezone\":\t\"+00:00\"\n\t}\n}") */
		var expected_ string = "Duplicate primary key `id`:\n{\n\t\"id\":\t{\n\t\t\"$reql_type$\":\t\"TIME\",\n\t\t\"epoch_time\":\t1375445163.087,\n\t\t\"timezone\":\t\"-07:00\"\n\t}\n}\n{\n\t\"id\":\t{\n\t\t\"$reql_type$\":\t\"TIME\",\n\t\t\"epoch_time\":\t1375445163.087,\n\t\t\"timezone\":\t\"+00:00\"\n\t}\n}"
		/* tbl.insert(bad_insert)['first_error'] */

		suite.T().Log("About to run line #42: tbl.Insert(bad_insert).AtIndex('first_error')")

		runAndAssert(suite.Suite, expected_, tbl.Insert(bad_insert).AtIndex("first_error"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #42")
	}

	{
		// times/index.yaml line #46
		/* 5 */
		var expected_ int = 5
		/* tbl.between(ts, te).count() */

		suite.T().Log("About to run line #46: tbl.Between(ts, te).Count()")

		runAndAssert(suite.Suite, expected_, tbl.Between(ts, te).Count(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #46")
	}

	{
		// times/index.yaml line #48
		/* 3 */
		var expected_ int = 3
		/* tbl.between(t1, t4).count() */

		suite.T().Log("About to run line #48: tbl.Between(t1, t4).Count()")

		runAndAssert(suite.Suite, expected_, tbl.Between(t1, t4).Count(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #48")
	}

	{
		// times/index.yaml line #51
		/* 4 */
		var expected_ int = 4
		/* tbl.between(t1, t4, right_bound='closed').count() */

		suite.T().Log("About to run line #51: tbl.Between(t1, t4).OptArgs(r.BetweenOpts{RightBound: 'closed', }).Count()")

		runAndAssert(suite.Suite, expected_, tbl.Between(t1, t4).OptArgs(r.BetweenOpts{RightBound: "closed"}).Count(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #51")
	}

	{
		// times/index.yaml line #54
		/* 5 */
		var expected_ int = 5
		/* tbl.between(r.expr(ts).in_timezone("+06:00"), te).count() */

		suite.T().Log("About to run line #54: tbl.Between(r.Expr(ts).InTimezone('+06:00'), te).Count()")

		runAndAssert(suite.Suite, expected_, tbl.Between(r.Expr(ts).InTimezone("+06:00"), te).Count(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #54")
	}

	{
		// times/index.yaml line #56
		/* 3 */
		var expected_ int = 3
		/* tbl.between(t1, r.expr(t4).in_timezone("+08:00")).count() */

		suite.T().Log("About to run line #56: tbl.Between(t1, r.Expr(t4).InTimezone('+08:00')).Count()")

		runAndAssert(suite.Suite, expected_, tbl.Between(t1, r.Expr(t4).InTimezone("+08:00")).Count(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #56")
	}

	{
		// times/index.yaml line #59
		/* 4 */
		var expected_ int = 4
		/* tbl.between(r.expr(t1).in_timezone("Z"), t4, right_bound='closed').count() */

		suite.T().Log("About to run line #59: tbl.Between(r.Expr(t1).InTimezone('Z'), t4).OptArgs(r.BetweenOpts{RightBound: 'closed', }).Count()")

		runAndAssert(suite.Suite, expected_, tbl.Between(r.Expr(t1).InTimezone("Z"), t4).OptArgs(r.BetweenOpts{RightBound: "closed"}).Count(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #59")
	}

	{
		// times/index.yaml line #64
		/* 5 */
		var expected_ int = 5
		/* tbl.update(lambda row:{'a':row['id']})['replaced'] */

		suite.T().Log("About to run line #64: tbl.Update(func(row r.Term) interface{} { return map[interface{}]interface{}{'a': row.AtIndex('id'), }}).AtIndex('replaced')")

		runAndAssert(suite.Suite, expected_, tbl.Update(func(row r.Term) interface{} { return map[interface{}]interface{}{"a": row.AtIndex("id")} }).AtIndex("replaced"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #64")
	}

	{
		// times/index.yaml line #67
		/* ({'created':1}) */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"created": 1}
		/* tbl.index_create('a') */

		suite.T().Log("About to run line #67: tbl.IndexCreate('a')")

		runAndAssert(suite.Suite, expected_, tbl.IndexCreate("a"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #67")
	}

	{
		// times/index.yaml line #69
		/* 1 */
		var expected_ int = 1
		/* tbl.index_wait('a').count() */

		suite.T().Log("About to run line #69: tbl.IndexWait('a').Count()")

		runAndAssert(suite.Suite, expected_, tbl.IndexWait("a").Count(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #69")
	}

	{
		// times/index.yaml line #73
		/* 5 */
		var expected_ int = 5
		/* tbl.between(ts, te, index='a').count() */

		suite.T().Log("About to run line #73: tbl.Between(ts, te).OptArgs(r.BetweenOpts{Index: 'a', }).Count()")

		runAndAssert(suite.Suite, expected_, tbl.Between(ts, te).OptArgs(r.BetweenOpts{Index: "a"}).Count(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #73")
	}

	{
		// times/index.yaml line #77
		/* 3 */
		var expected_ int = 3
		/* tbl.between(t1, t4, index='a').count() */

		suite.T().Log("About to run line #77: tbl.Between(t1, t4).OptArgs(r.BetweenOpts{Index: 'a', }).Count()")

		runAndAssert(suite.Suite, expected_, tbl.Between(t1, t4).OptArgs(r.BetweenOpts{Index: "a"}).Count(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #77")
	}

	{
		// times/index.yaml line #81
		/* 4 */
		var expected_ int = 4
		/* tbl.between(t1, t4, right_bound='closed', index='a').count() */

		suite.T().Log("About to run line #81: tbl.Between(t1, t4).OptArgs(r.BetweenOpts{RightBound: 'closed', Index: 'a', }).Count()")

		runAndAssert(suite.Suite, expected_, tbl.Between(t1, t4).OptArgs(r.BetweenOpts{RightBound: "closed", Index: "a"}).Count(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #81")
	}

	{
		// times/index.yaml line #85
		/* 5 */
		var expected_ int = 5
		/* tbl.between(r.expr(ts).in_timezone("+06:00"), te, index='a').count() */

		suite.T().Log("About to run line #85: tbl.Between(r.Expr(ts).InTimezone('+06:00'), te).OptArgs(r.BetweenOpts{Index: 'a', }).Count()")

		runAndAssert(suite.Suite, expected_, tbl.Between(r.Expr(ts).InTimezone("+06:00"), te).OptArgs(r.BetweenOpts{Index: "a"}).Count(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #85")
	}

	{
		// times/index.yaml line #89
		/* 3 */
		var expected_ int = 3
		/* tbl.between(t1, r.expr(t4).in_timezone("+08:00"), index='a').count() */

		suite.T().Log("About to run line #89: tbl.Between(t1, r.Expr(t4).InTimezone('+08:00')).OptArgs(r.BetweenOpts{Index: 'a', }).Count()")

		runAndAssert(suite.Suite, expected_, tbl.Between(t1, r.Expr(t4).InTimezone("+08:00")).OptArgs(r.BetweenOpts{Index: "a"}).Count(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #89")
	}

	{
		// times/index.yaml line #93
		/* 4 */
		var expected_ int = 4
		/* tbl.between(r.expr(t1).in_timezone("Z"), t4, right_bound='closed', index='a').count() */

		suite.T().Log("About to run line #93: tbl.Between(r.Expr(t1).InTimezone('Z'), t4).OptArgs(r.BetweenOpts{RightBound: 'closed', Index: 'a', }).Count()")

		runAndAssert(suite.Suite, expected_, tbl.Between(r.Expr(t1).InTimezone("Z"), t4).OptArgs(r.BetweenOpts{RightBound: "closed", Index: "a"}).Count(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #93")
	}

	{
		// times/index.yaml line #98
		/* ({'created':1}) */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"created": 1}
		/* tbl.index_create('b', lambda row:r.branch(row['id'] < t4, row['a'], null)) */

		suite.T().Log("About to run line #98: tbl.IndexCreateFunc('b', func(row r.Term) interface{} { return r.Branch(row.AtIndex('id').Lt(t4), row.AtIndex('a'), nil)})")

		runAndAssert(suite.Suite, expected_, tbl.IndexCreateFunc("b", func(row r.Term) interface{} { return r.Branch(row.AtIndex("id").Lt(t4), row.AtIndex("a"), nil) }), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #98")
	}

	{
		// times/index.yaml line #101
		/* 1 */
		var expected_ int = 1
		/* tbl.index_wait('b').count() */

		suite.T().Log("About to run line #101: tbl.IndexWait('b').Count()")

		runAndAssert(suite.Suite, expected_, tbl.IndexWait("b").Count(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #101")
	}

	{
		// times/index.yaml line #105
		/* 1 */
		var expected_ int = 1
		/* tbl.index_wait('b').count() */

		suite.T().Log("About to run line #105: tbl.IndexWait('b').Count()")

		runAndAssert(suite.Suite, expected_, tbl.IndexWait("b").Count(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #105")
	}

	{
		// times/index.yaml line #109
		/* 3 */
		var expected_ int = 3
		/* tbl.between(ts, te, index='b').count() */

		suite.T().Log("About to run line #109: tbl.Between(ts, te).OptArgs(r.BetweenOpts{Index: 'b', }).Count()")

		runAndAssert(suite.Suite, expected_, tbl.Between(ts, te).OptArgs(r.BetweenOpts{Index: "b"}).Count(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #109")
	}

	{
		// times/index.yaml line #113
		/* 3 */
		var expected_ int = 3
		/* tbl.between(t1, t4, index='b').count() */

		suite.T().Log("About to run line #113: tbl.Between(t1, t4).OptArgs(r.BetweenOpts{Index: 'b', }).Count()")

		runAndAssert(suite.Suite, expected_, tbl.Between(t1, t4).OptArgs(r.BetweenOpts{Index: "b"}).Count(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #113")
	}

	{
		// times/index.yaml line #117
		/* 3 */
		var expected_ int = 3
		/* tbl.between(t1, t4, right_bound='closed', index='b').count() */

		suite.T().Log("About to run line #117: tbl.Between(t1, t4).OptArgs(r.BetweenOpts{RightBound: 'closed', Index: 'b', }).Count()")

		runAndAssert(suite.Suite, expected_, tbl.Between(t1, t4).OptArgs(r.BetweenOpts{RightBound: "closed", Index: "b"}).Count(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #117")
	}

	{
		// times/index.yaml line #121
		/* 3 */
		var expected_ int = 3
		/* tbl.between(r.expr(ts).in_timezone("+06:00"), te, index='b').count() */

		suite.T().Log("About to run line #121: tbl.Between(r.Expr(ts).InTimezone('+06:00'), te).OptArgs(r.BetweenOpts{Index: 'b', }).Count()")

		runAndAssert(suite.Suite, expected_, tbl.Between(r.Expr(ts).InTimezone("+06:00"), te).OptArgs(r.BetweenOpts{Index: "b"}).Count(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #121")
	}

	{
		// times/index.yaml line #125
		/* 3 */
		var expected_ int = 3
		/* tbl.between(t1, r.expr(t4).in_timezone("+08:00"), index='b').count() */

		suite.T().Log("About to run line #125: tbl.Between(t1, r.Expr(t4).InTimezone('+08:00')).OptArgs(r.BetweenOpts{Index: 'b', }).Count()")

		runAndAssert(suite.Suite, expected_, tbl.Between(t1, r.Expr(t4).InTimezone("+08:00")).OptArgs(r.BetweenOpts{Index: "b"}).Count(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #125")
	}

	{
		// times/index.yaml line #129
		/* 3 */
		var expected_ int = 3
		/* tbl.between(r.expr(t1).in_timezone("Z"), t4, right_bound='closed', index='b').count() */

		suite.T().Log("About to run line #129: tbl.Between(r.Expr(t1).InTimezone('Z'), t4).OptArgs(r.BetweenOpts{RightBound: 'closed', Index: 'b', }).Count()")

		runAndAssert(suite.Suite, expected_, tbl.Between(r.Expr(t1).InTimezone("Z"), t4).OptArgs(r.BetweenOpts{RightBound: "closed", Index: "b"}).Count(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #129")
	}

	// times/index.yaml line #135
	// oldtime = datetime.fromtimestamp(1375147296.681, PacificTimeZone())
	suite.T().Log("Possibly executing: var oldtime time.Time = Ast.Fromtimestamp(1375147296.681, PacificTimeZone())")

	oldtime := Ast.Fromtimestamp(1375147296.681, PacificTimeZone())
	_ = oldtime // Prevent any noused variable errors

	// times/index.yaml line #139
	// curtime = datetime.now()
	suite.T().Log("Possibly executing: var curtime time.Time = Ast.Now()")

	curtime := Ast.Now()
	_ = curtime // Prevent any noused variable errors

	{
		// times/index.yaml line #142
		/* 1 */
		var expected_ int = 1
		/* tbl.insert([{'id':oldtime}])['inserted'] */

		suite.T().Log("About to run line #142: tbl.Insert([]interface{}{map[interface{}]interface{}{'id': oldtime, }}).AtIndex('inserted')")

		runAndAssert(suite.Suite, expected_, tbl.Insert([]interface{}{map[interface{}]interface{}{"id": oldtime}}).AtIndex("inserted"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #142")
	}

	{
		// times/index.yaml line #148
		/* ("PTYPE<TIME>") */
		var expected_ string = "PTYPE<TIME>"
		/* tbl.get(oldtime)['id'].type_of() */

		suite.T().Log("About to run line #148: tbl.Get(oldtime).AtIndex('id').TypeOf()")

		runAndAssert(suite.Suite, expected_, tbl.Get(oldtime).AtIndex("id").TypeOf(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #148")
	}
}
