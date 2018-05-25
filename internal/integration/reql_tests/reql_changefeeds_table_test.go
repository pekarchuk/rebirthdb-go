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

// Test changefeeds on a table
func TestChangefeedsTableSuite(t *testing.T) {
	suite.Run(t, new(ChangefeedsTableSuite))
}

type ChangefeedsTableSuite struct {
	suite.Suite

	session *r.Session
}

func (suite *ChangefeedsTableSuite) SetupTest() {
	suite.T().Log("Setting up ChangefeedsTableSuite")
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

func (suite *ChangefeedsTableSuite) TearDownSuite() {
	suite.T().Log("Tearing down ChangefeedsTableSuite")

	if suite.session != nil {
		r.DB("rethinkdb").Table("_debug_scratch").Delete().Exec(suite.session)
		r.DB("test").TableDrop("tbl").Exec(suite.session)
		r.DBDrop("test").Exec(suite.session)

		suite.session.Close()
	}
}

func (suite *ChangefeedsTableSuite) TestCases() {
	suite.T().Log("Running ChangefeedsTableSuite: Test changefeeds on a table")

	tbl := r.DB("test").Table("tbl")
	_ = tbl // Prevent any noused variable errors

	// changefeeds/table.yaml line #9
	// all = tbl.changes()
	suite.T().Log("Possibly executing: var all r.Term = tbl.Changes()")

	all := maybeRun(tbl.Changes(), suite.session, r.RunOpts{})
	_ = all // Prevent any noused variable errors

	{
		// changefeeds/table.yaml line #15
		/* partial({'errors':0, 'inserted':2}) */
		var expected_ compare.Expected = compare.PartialMatch(map[interface{}]interface{}{"errors": 0, "inserted": 2})
		/* tbl.insert([{'id':1}, {'id':2}]) */

		suite.T().Log("About to run line #15: tbl.Insert([]interface{}{map[interface{}]interface{}{'id': 1, }, map[interface{}]interface{}{'id': 2, }})")

		runAndAssert(suite.Suite, expected_, tbl.Insert([]interface{}{map[interface{}]interface{}{"id": 1}, map[interface{}]interface{}{"id": 2}}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #15")
	}

	{
		// changefeeds/table.yaml line #17
		/* bag([{'old_val':null, 'new_val':{'id':1}}, {'old_val':null, 'new_val':{'id':2}}]) */
		var expected_ compare.Expected = compare.UnorderedMatch([]interface{}{map[interface{}]interface{}{"old_val": nil, "new_val": map[interface{}]interface{}{"id": 1}}, map[interface{}]interface{}{"old_val": nil, "new_val": map[interface{}]interface{}{"id": 2}}})
		/* fetch(all, 2) */

		suite.T().Log("About to run line #17: fetch(all, 2)")

		fetchAndAssert(suite.Suite, expected_, all, 2)
		suite.T().Log("Finished running line #17")
	}

	{
		// changefeeds/table.yaml line #22
		/* partial({'errors':0, 'replaced':1}) */
		var expected_ compare.Expected = compare.PartialMatch(map[interface{}]interface{}{"errors": 0, "replaced": 1})
		/* tbl.get(1).update({'version':1}) */

		suite.T().Log("About to run line #22: tbl.Get(1).Update(map[interface{}]interface{}{'version': 1, })")

		runAndAssert(suite.Suite, expected_, tbl.Get(1).Update(map[interface{}]interface{}{"version": 1}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #22")
	}

	{
		// changefeeds/table.yaml line #24
		/* [{'old_val':{'id':1}, 'new_val':{'id':1, 'version':1}}] */
		var expected_ []interface{} = []interface{}{map[interface{}]interface{}{"old_val": map[interface{}]interface{}{"id": 1}, "new_val": map[interface{}]interface{}{"id": 1, "version": 1}}}
		/* fetch(all, 1) */

		suite.T().Log("About to run line #24: fetch(all, 1)")

		fetchAndAssert(suite.Suite, expected_, all, 1)
		suite.T().Log("Finished running line #24")
	}

	{
		// changefeeds/table.yaml line #29
		/* partial({'errors':0, 'deleted':1}) */
		var expected_ compare.Expected = compare.PartialMatch(map[interface{}]interface{}{"errors": 0, "deleted": 1})
		/* tbl.get(1).delete() */

		suite.T().Log("About to run line #29: tbl.Get(1).Delete()")

		runAndAssert(suite.Suite, expected_, tbl.Get(1).Delete(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #29")
	}

	{
		// changefeeds/table.yaml line #31
		/* [{'old_val':{'id':1, 'version':1}, 'new_val':null}] */
		var expected_ []interface{} = []interface{}{map[interface{}]interface{}{"old_val": map[interface{}]interface{}{"id": 1, "version": 1}, "new_val": nil}}
		/* fetch(all, 1) */

		suite.T().Log("About to run line #31: fetch(all, 1)")

		fetchAndAssert(suite.Suite, expected_, all, 1)
		suite.T().Log("Finished running line #31")
	}

	// changefeeds/table.yaml line #36
	// pluck = tbl.changes().pluck({'new_val':['version']})
	suite.T().Log("Possibly executing: var pluck r.Term = tbl.Changes().Pluck(map[interface{}]interface{}{'new_val': []interface{}{'version'}, })")

	pluck := maybeRun(tbl.Changes().Pluck(map[interface{}]interface{}{"new_val": []interface{}{"version"}}), suite.session, r.RunOpts{})
	_ = pluck // Prevent any noused variable errors

	{
		// changefeeds/table.yaml line #37
		/* partial({'errors':0, 'inserted':1}) */
		var expected_ compare.Expected = compare.PartialMatch(map[interface{}]interface{}{"errors": 0, "inserted": 1})
		/* tbl.insert([{'id':5, 'version':5}]) */

		suite.T().Log("About to run line #37: tbl.Insert([]interface{}{map[interface{}]interface{}{'id': 5, 'version': 5, }})")

		runAndAssert(suite.Suite, expected_, tbl.Insert([]interface{}{map[interface{}]interface{}{"id": 5, "version": 5}}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #37")
	}

	{
		// changefeeds/table.yaml line #39
		/* [{'new_val':{'version':5}}] */
		var expected_ []interface{} = []interface{}{map[interface{}]interface{}{"new_val": map[interface{}]interface{}{"version": 5}}}
		/* fetch(pluck, 1) */

		suite.T().Log("About to run line #39: fetch(pluck, 1)")

		fetchAndAssert(suite.Suite, expected_, pluck, 1)
		suite.T().Log("Finished running line #39")
	}

	{
		// changefeeds/table.yaml line #44
		/* err('ReqlQueryLogicError', "Cannot call a terminal (`reduce`, `count`, etc.) on an infinite stream (such as a changefeed).") */
		var expected_ Err = err("ReqlQueryLogicError", "Cannot call a terminal (`reduce`, `count`, etc.) on an infinite stream (such as a changefeed).")
		/* tbl.changes().order_by('id') */

		suite.T().Log("About to run line #44: tbl.Changes().OrderBy('id')")

		runAndAssert(suite.Suite, expected_, tbl.Changes().OrderBy("id"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #44")
	}

	// changefeeds/table.yaml line #59
	// overflow = tbl.changes(changefeed_queue_size=100)
	suite.T().Log("Possibly executing: var overflow r.Term = tbl.Changes().OptArgs(r.ChangesOpts{ChangefeedQueueSize: 100, })")

	overflow := maybeRun(tbl.Changes().OptArgs(r.ChangesOpts{ChangefeedQueueSize: 100}), suite.session, r.RunOpts{})
	_ = overflow // Prevent any noused variable errors

	{
		// changefeeds/table.yaml line #62
		/* AnythingIsFine */
		var expected_ string = compare.AnythingIsFine
		/* tbl.insert(r.range(200).map(lambda x: {})) */

		suite.T().Log("About to run line #62: tbl.Insert(r.Range(200).Map(func(x r.Term) interface{} { return map[interface{}]interface{}{}}))")

		runAndAssert(suite.Suite, expected_, tbl.Insert(r.Range(200).Map(func(x r.Term) interface{} { return map[interface{}]interface{}{} })), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #62")
	}

	{
		// changefeeds/table.yaml line #64
		/* partial([{'error': regex('Changefeed cache over array size limit, skipped \d+ elements.')}]) */
		var expected_ compare.Expected = compare.PartialMatch([]interface{}{map[interface{}]interface{}{"error": compare.MatchesRegexp("Changefeed cache over array size limit, skipped \\d+ elements.")}})
		/* fetch(overflow, 90) */

		suite.T().Log("About to run line #64: fetch(overflow, 90)")

		fetchAndAssert(suite.Suite, expected_, overflow, 90)
		suite.T().Log("Finished running line #64")
	}

	// changefeeds/table.yaml line #69
	// vtbl = r.db('rethinkdb').table('_debug_scratch')
	suite.T().Log("Possibly executing: var vtbl r.Term = r.DB('rethinkdb').Table('_debug_scratch')")

	vtbl := r.DB("rethinkdb").Table("_debug_scratch")
	_ = vtbl // Prevent any noused variable errors

	// changefeeds/table.yaml line #70
	// allVirtual = vtbl.changes()
	suite.T().Log("Possibly executing: var allVirtual r.Term = vtbl.Changes()")

	allVirtual := maybeRun(vtbl.Changes(), suite.session, r.RunOpts{})
	_ = allVirtual // Prevent any noused variable errors

	{
		// changefeeds/table.yaml line #74
		/* partial({'errors':0, 'inserted':2}) */
		var expected_ compare.Expected = compare.PartialMatch(map[interface{}]interface{}{"errors": 0, "inserted": 2})
		/* vtbl.insert([{'id':1}, {'id':2}]) */

		suite.T().Log("About to run line #74: vtbl.Insert([]interface{}{map[interface{}]interface{}{'id': 1, }, map[interface{}]interface{}{'id': 2, }})")

		runAndAssert(suite.Suite, expected_, vtbl.Insert([]interface{}{map[interface{}]interface{}{"id": 1}, map[interface{}]interface{}{"id": 2}}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #74")
	}

	{
		// changefeeds/table.yaml line #76
		/* bag([{'old_val':null, 'new_val':{'id':1}}, {'old_val':null, 'new_val':{'id':2}}]) */
		var expected_ compare.Expected = compare.UnorderedMatch([]interface{}{map[interface{}]interface{}{"old_val": nil, "new_val": map[interface{}]interface{}{"id": 1}}, map[interface{}]interface{}{"old_val": nil, "new_val": map[interface{}]interface{}{"id": 2}}})
		/* fetch(allVirtual, 2) */

		suite.T().Log("About to run line #76: fetch(allVirtual, 2)")

		fetchAndAssert(suite.Suite, expected_, allVirtual, 2)
		suite.T().Log("Finished running line #76")
	}

	{
		// changefeeds/table.yaml line #81
		/* partial({'errors':0, 'replaced':1}) */
		var expected_ compare.Expected = compare.PartialMatch(map[interface{}]interface{}{"errors": 0, "replaced": 1})
		/* vtbl.get(1).update({'version':1}) */

		suite.T().Log("About to run line #81: vtbl.Get(1).Update(map[interface{}]interface{}{'version': 1, })")

		runAndAssert(suite.Suite, expected_, vtbl.Get(1).Update(map[interface{}]interface{}{"version": 1}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #81")
	}

	{
		// changefeeds/table.yaml line #83
		/* [{'old_val':{'id':1}, 'new_val':{'id':1, 'version':1}}] */
		var expected_ []interface{} = []interface{}{map[interface{}]interface{}{"old_val": map[interface{}]interface{}{"id": 1}, "new_val": map[interface{}]interface{}{"id": 1, "version": 1}}}
		/* fetch(allVirtual, 1) */

		suite.T().Log("About to run line #83: fetch(allVirtual, 1)")

		fetchAndAssert(suite.Suite, expected_, allVirtual, 1)
		suite.T().Log("Finished running line #83")
	}

	{
		// changefeeds/table.yaml line #88
		/* partial({'errors':0, 'deleted':1}) */
		var expected_ compare.Expected = compare.PartialMatch(map[interface{}]interface{}{"errors": 0, "deleted": 1})
		/* vtbl.get(1).delete() */

		suite.T().Log("About to run line #88: vtbl.Get(1).Delete()")

		runAndAssert(suite.Suite, expected_, vtbl.Get(1).Delete(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #88")
	}

	{
		// changefeeds/table.yaml line #90
		/* [{'old_val':{'id':1, 'version':1}, 'new_val':null}] */
		var expected_ []interface{} = []interface{}{map[interface{}]interface{}{"old_val": map[interface{}]interface{}{"id": 1, "version": 1}, "new_val": nil}}
		/* fetch(allVirtual, 1) */

		suite.T().Log("About to run line #90: fetch(allVirtual, 1)")

		fetchAndAssert(suite.Suite, expected_, allVirtual, 1)
		suite.T().Log("Finished running line #90")
	}

	// changefeeds/table.yaml line #95
	// vpluck = vtbl.changes().pluck({'new_val':['version']})
	suite.T().Log("Possibly executing: var vpluck r.Term = vtbl.Changes().Pluck(map[interface{}]interface{}{'new_val': []interface{}{'version'}, })")

	vpluck := maybeRun(vtbl.Changes().Pluck(map[interface{}]interface{}{"new_val": []interface{}{"version"}}), suite.session, r.RunOpts{})
	_ = vpluck // Prevent any noused variable errors

	{
		// changefeeds/table.yaml line #96
		/* partial({'errors':0, 'inserted':1}) */
		var expected_ compare.Expected = compare.PartialMatch(map[interface{}]interface{}{"errors": 0, "inserted": 1})
		/* vtbl.insert([{'id':5, 'version':5}]) */

		suite.T().Log("About to run line #96: vtbl.Insert([]interface{}{map[interface{}]interface{}{'id': 5, 'version': 5, }})")

		runAndAssert(suite.Suite, expected_, vtbl.Insert([]interface{}{map[interface{}]interface{}{"id": 5, "version": 5}}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #96")
	}

	{
		// changefeeds/table.yaml line #98
		/* [{'new_val':{'version':5}}] */
		var expected_ []interface{} = []interface{}{map[interface{}]interface{}{"new_val": map[interface{}]interface{}{"version": 5}}}
		/* fetch(vpluck, 1) */

		suite.T().Log("About to run line #98: fetch(vpluck, 1)")

		fetchAndAssert(suite.Suite, expected_, vpluck, 1)
		suite.T().Log("Finished running line #98")
	}
}
