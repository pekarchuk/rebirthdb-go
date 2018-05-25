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

// Tests that manipulation data in tables
func TestPolymorphismSuite(t *testing.T) {
	suite.Run(t, new(PolymorphismSuite))
}

type PolymorphismSuite struct {
	suite.Suite

	session *r.Session
}

func (suite *PolymorphismSuite) SetupTest() {
	suite.T().Log("Setting up PolymorphismSuite")
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

func (suite *PolymorphismSuite) TearDownSuite() {
	suite.T().Log("Tearing down PolymorphismSuite")

	if suite.session != nil {
		r.DB("rethinkdb").Table("_debug_scratch").Delete().Exec(suite.session)
		r.DB("test").TableDrop("tbl").Exec(suite.session)
		r.DBDrop("test").Exec(suite.session)

		suite.session.Close()
	}
}

func (suite *PolymorphismSuite) TestCases() {
	suite.T().Log("Running PolymorphismSuite: Tests that manipulation data in tables")

	tbl := r.DB("test").Table("tbl")
	_ = tbl // Prevent any noused variable errors

	// polymorphism.yaml line #5
	// obj = r.expr({'id':0,'a':0})
	suite.T().Log("Possibly executing: var obj r.Term = r.Expr(map[interface{}]interface{}{'id': 0, 'a': 0, })")

	obj := r.Expr(map[interface{}]interface{}{"id": 0, "a": 0})
	_ = obj // Prevent any noused variable errors

	{
		// polymorphism.yaml line #7
		/* ({'deleted':0,'replaced':0,'unchanged':0,'errors':0,'skipped':0,'inserted':3}) */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"deleted": 0, "replaced": 0, "unchanged": 0, "errors": 0, "skipped": 0, "inserted": 3}
		/* tbl.insert([{'id':i, 'a':i} for i in xrange(3)]) */

		suite.T().Log("About to run line #7: tbl.Insert((func() []interface{} {\n    res := []interface{}{}\n    for iterator_ := 0; iterator_ < 3; iterator_++ {\n        i := iterator_\n        res = append(res, map[interface{}]interface{}{'id': i, 'a': i, })\n    }\n    return res\n}()))")

		runAndAssert(suite.Suite, expected_, tbl.Insert((func() []interface{} {
			res := []interface{}{}
			for iterator_ := 0; iterator_ < 3; iterator_++ {
				i := iterator_
				res = append(res, map[interface{}]interface{}{"id": i, "a": i})
			}
			return res
		}())), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #7")
	}

	{
		// polymorphism.yaml line #21
		/* ({'id':0,'c':1,'a':0}) */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"id": 0, "c": 1, "a": 0}
		/* tbl.merge({'c':1}).nth(0) */

		suite.T().Log("About to run line #21: tbl.Merge(map[interface{}]interface{}{'c': 1, }).Nth(0)")

		runAndAssert(suite.Suite, expected_, tbl.Merge(map[interface{}]interface{}{"c": 1}).Nth(0), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #21")
	}

	{
		// polymorphism.yaml line #22
		/* ({'id':0,'c':1,'a':0}) */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"id": 0, "c": 1, "a": 0}
		/* obj.merge({'c':1}) */

		suite.T().Log("About to run line #22: obj.Merge(map[interface{}]interface{}{'c': 1, })")

		runAndAssert(suite.Suite, expected_, obj.Merge(map[interface{}]interface{}{"c": 1}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #22")
	}

	{
		// polymorphism.yaml line #26
		/* ({'id':0}) */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"id": 0}
		/* tbl.without('a').nth(0) */

		suite.T().Log("About to run line #26: tbl.Without('a').Nth(0)")

		runAndAssert(suite.Suite, expected_, tbl.Without("a").Nth(0), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #26")
	}

	{
		// polymorphism.yaml line #27
		/* ({'id':0}) */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"id": 0}
		/* obj.without('a') */

		suite.T().Log("About to run line #27: obj.Without('a')")

		runAndAssert(suite.Suite, expected_, obj.Without("a"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #27")
	}

	{
		// polymorphism.yaml line #31
		/* ({'a':0}) */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"a": 0}
		/* tbl.pluck('a').nth(0) */

		suite.T().Log("About to run line #31: tbl.Pluck('a').Nth(0)")

		runAndAssert(suite.Suite, expected_, tbl.Pluck("a").Nth(0), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #31")
	}

	{
		// polymorphism.yaml line #32
		/* ({'a':0}) */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"a": 0}
		/* obj.pluck('a') */

		suite.T().Log("About to run line #32: obj.Pluck('a')")

		runAndAssert(suite.Suite, expected_, obj.Pluck("a"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #32")
	}
}
