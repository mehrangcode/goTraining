import axios from "axios";
import { ConstValues } from "utils/const";


export async function getAll() {
    return await axios.get(ConstValues + "/users")
}

export async function create(payload) {
    return await axios.post(ConstValues + "/users", payload)
}