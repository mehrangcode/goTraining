import { ConstValues } from "@src/utils/const";
import axios from "axios";

const api = ConstValues.backUri + "/api/tables"
export async function getAll() {
    return await axios.get(api)
}

export async function createTable(payload) {
    return await axios.post(api, payload)
}
export async function updateTable(tableId, payload) {
    return await axios.put(api + "/" + tableId, payload)
}

export async function deleteTable(tableId) {
    return await axios.delete(api + "/" + tableId)
}