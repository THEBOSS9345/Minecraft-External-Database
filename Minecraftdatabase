import * as net from "@minecraft/server-net";

export default   class Database {
    constructor(fileName) {
        this.baseURL = "http://localhost:8080";
        this.fileName = fileName;
    }

   async set(key, value) {
        return new Promise((resolve, reject) => {
            const req = new net.HttpRequest(`${this.baseURL}/set?file=${encodeURIComponent(this.fileName)}`);
            req.setBody(JSON.stringify({[key]: value}));
            req.method = net.HttpRequestMethod.Post
            req.addHeader("Content-Type", "application/json");

            net.http.request(req).then((res) => {
                if (res.status === 200) {
                    return resolve(true);
                } else {
                    const error = JSON.parse(res.body);
                    return reject(new Error(`Error setting key: ${JSON.stringify(error)}`));
                }
            });
        })
    }


    async get(key) {
        return new Promise((resolve, reject) => {
            const req = new net.HttpRequest(`${this.baseURL}/get?file=${encodeURIComponent(this.fileName)}&key=${encodeURIComponent(key)}`);
           req.method = net.HttpRequestMethod.Get;
            net.http.request(req).then((res) => {
                if (res.status === 200) {
                    const data = JSON.parse(res.body);
                    return resolve(data);
                } else {
                    const error = JSON.parse(res.body);
                    return reject(new Error(`Error getting key: ${JSON.stringify(error)}`));
                }
            });
        })
    }

    async delete(key) {
        return new Promise((resolve, reject) => {
            const req = new net.HttpRequest(`${this.baseURL}/delete?file=${encodeURIComponent(this.fileName)}&key=${encodeURIComponent(key)}`);
            req.method = net.HttpRequestMethod.Delete
            net.http.request(req).then((res) => {
                if (res.status === 200) {
                    return resolve(`Key "${key}" deleted successfully`);
                } else {
                    const error = JSON.parse(res.body);
                    return reject(new Error(`Error deleting key: ${JSON.stringify(error)}`));
                }
            });
        })
    }

    async has(key) {
        return new Promise((resolve, reject) => {
            const req = new net.HttpRequest(`${this.baseURL}/has?file=${encodeURIComponent(this.fileName)}&key=${encodeURIComponent(key)}`);
            req.method = net.HttpRequestMethod.Get
            net.http.request(req).then((res) => {
                if (res.status === 200) {
                    const exists = JSON.parse(res.body);
                    return resolve(exists?.exists || false);
                } else {
                    const error = JSON.parse(res.body);
                    return reject(new Error(`Error checking key existence: ${error.message}`));
                }
            });
        })
    }

    async list() {
        return new Promise((resolve, reject) => {
            const req = new net.HttpRequest(`${this.baseURL}/list?file=${encodeURIComponent(this.fileName)}`);
            req.method = net.HttpRequestMethod.Get
            net.http.request(req).then((res) => {
                if (res.status === 200) {
                    const keys = JSON.parse(res.body);
                    return resolve(keys);
                } else {
                    const error = JSON.parse(res.body);
                    return reject(new Error(`Error listing keys: ${JSON.stringify(error)}`));
                }
            });
        })
    }
}
