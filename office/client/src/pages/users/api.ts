import { ConstValues } from "@src/utils/const";
import axios from "axios";


export async function getAll() {
    return await axios.get(ConstValues.backUri + "/users")
}

export async function create(payload) {
    return await axios.post(ConstValues.backUri + "/users", payload)
}