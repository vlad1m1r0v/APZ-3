import requests from "../common/http.js";

class Client {
	constructor(url) {
		this.requests = new requests(url);
	}

	getMachines() {
		this.requests.get("/machines");
	}

	getDisks() {
		this.requests.get("/disks");
	}

	postMachine(data) {
		this.requests.post("/machines", data);
	}

	postDisk(data) {
		this.requests.post("/disk", data);
	}
}

export default Client;
