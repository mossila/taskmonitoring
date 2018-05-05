var chakram = require('chakram');
var expect = chakram.expect;
var base = "http://localhost:8080/";

describe("task/", function () {
    it("should provide list of existing tasks", function () {
        var response = chakram.get(base + "/task")
        expect(response).to.have.status(200);
        expect(response).to.have.schema({
            type: "array"
        })
        return chakram.wait();
    });

    it("should provide an item when given an existing id", function () {
        var response = chakram.get(base + "task/1")
        expect(response).to.have.status(200);
        expect(response).to.have.schema(
            {
                "type": "object",
                properties: {
                    id: { type: "number" },
                    name: { type: "string" },
                    completed: { type: "boolean" },
                    due: { type: "string" }
                }
            }
        )
        return chakram.wait();
    });

    it("should create new item when post with new item", function () {
        body = {
            name: "test sending a string",
            completed: false,
            due: new Date()
        };
        var response = chakram.post(base + "task", body);
        expect(response).to.have.status(200);
        expect(response).to.have.schema(
            {
                "type": "object",
                properties: {
                    id: { type: "number" },
                    name: { type: "string" },
                    completed: { type: "boolean" },
                    due: { type: "string" }
                }
            }
        )
        return chakram.wait();
    })
});
