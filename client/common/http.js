import axios from "axios";

class requests {
	constructor(baseUrl) {
		this.baseUrl = baseUrl;
	}

	get(path) {
		return axios({
			method: "get",
			url: `${this.baseUrl}${path}`
		});
	}

	post(path, data) {
		return axios({
			method: "post",
			url: `${this.baseUrl}${path}`,
			data: data
		});
	}

	put(path, data) {
		return fetch(`${this.baseUrl}${path}`, {
			method: "put",
			data: data
		});
	}

	delete(path) {
		return fetch({
			method: "delete",
			url: `${this.baseUrl}${path}`
		});
	}
}

export default requests;
