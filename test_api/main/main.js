var chakram = require('chakram');
var expect = chakram.expect;
describe("task/", function() {
    it("should provide list of existing tasks", function() {
        var response =  chakram.get("http://localhost:8080/task")
        return expect(response).to.have.status(200);
    });
});