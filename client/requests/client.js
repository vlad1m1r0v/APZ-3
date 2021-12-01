import requests from "../common/http.js";

class Client {
	constructor(url) {
		this.requests = new requests(url);
	}

	getMachines() {
		return this.requests.get("/machines");
	}

	connectToMachine(id, machineId) {
		return this.requests.put(`/disks/${id}`, { machineId });
	}
}

export default Client;
