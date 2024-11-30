const baseURL = 'http://localhost:8080';

class DatabaseManager {
    constructor(file) {
        if (!file) {
            throw new Error('A database file name must be provided.');
        }
        this.baseURL = baseURL;
        this.file = file;
    }

    async set(key, value) {
        try {
            const response = await fetch(`${this.baseURL}/set?file=${encodeURIComponent(this.file)}`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({ [key]: value }),
            });

            if (!response.ok) {
                const error = await response.text();
                throw new Error(`Failed to set data: ${error}`);
            }

            return true
        } catch (error) {
            console.error('Error in set:', error);
            return error
        }
    }

    async get(key) {
        try {
            const response = await fetch(`${this.baseURL}/get?file=${encodeURIComponent(this.file)}&key=${encodeURIComponent(key)}`);
            if (!response.ok) {
                const error = await response.text();
                throw new Error(`Failed to get data: ${error}`);
            }

            return (await response.json())
        } catch (error) {
            console.error('Error in get:', error);
            return error
        }
    }

    async delete(key) {
        try {
            const response = await fetch(`${this.baseURL}/delete?file=${encodeURIComponent(this.file)}&key=${encodeURIComponent(key)}`, {
                method: 'DELETE',
            });

            if (!response.ok) {
                const error = await response.text();
                throw new Error(`Failed to delete data: ${error}`);
            }

            return true
        } catch (error) {
            console.error('Error in delete:', error);
            return error
        }
    }

    async has(key) {
        try {
            const response = await fetch(`${this.baseURL}/has?file=${encodeURIComponent(this.file)}&key=${encodeURIComponent(key)}`);
            if (!response.ok) {
                const error = await response.text();
                throw new Error(`Failed to check data: ${error}`);
            }

            return (await response.json())?.exists;
        } catch (error) {
            console.error('Error in has:', error);
            return error
        }
    }

    async list() {
        try {
            const response = await fetch(`${this.baseURL}/list?file=${encodeURIComponent(this.file)}`);
            if (!response.ok) {
                const error = await response.text();
                throw new Error(`Failed to list keys: ${error}`);
            }

            return (await response.json());
        } catch (error) {
            console.error('Error in list:', error);
            return error
        }
    }
}