import { ConstValues } from "@src/utils/const";
import axios from "axios";

const api = ConstValues.backUri + "/api/users"
export async function getAll() {
    return await axios.get(api)
}

export async function createUser(payload) {
    return await axios.post(api, payload)
}
export async function updateUser(userId, payload) {
    return await axios.put(api + "/" + userId, payload)
}

export async function deleteUser(userId) {
    return await axios.delete(api + "/" + userId)
}