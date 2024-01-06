import { ConstValues } from "@src/utils/const";
import axios from "axios";


export async function getAll() {
    return await axios.get(ConstValues.backUri + "/users")
}

export async function createUser(payload) {
    return await axios.post(ConstValues.backUri + "/users", payload)
}
export async function updateUser(userId, payload) {
    return await axios.put(ConstValues.backUri + "/users/" + userId, payload)
}

export async function deleteUser(userId) {
    return await axios.delete(ConstValues.backUri + "/users/" + userId)
}