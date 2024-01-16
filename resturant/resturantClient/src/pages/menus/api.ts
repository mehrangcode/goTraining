import { ConstValues } from "@src/utils/const";
import axios from "axios";

const api = ConstValues.backUri + "/api/menus"
export async function getAll() {
    return await axios.get(api)
}

export async function createMenu(payload) {
    return await axios.post(api, payload)
}
export async function updateMenu(menuId, payload) {
    return await axios.put(api + "/" + menuId, payload)
}

export async function deleteMenu(menuId) {
    return await axios.delete(api + "/" + menuId)
}