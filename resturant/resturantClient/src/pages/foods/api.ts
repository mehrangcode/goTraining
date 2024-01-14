import { ConstValues } from "@src/utils/const";
import axios from "axios";

const api = ConstValues.backUri + "/api/foods"
export async function getAll() {
    return await axios.get(api)
}
export async function createFood(payload) {
    return await axios.post(api, payload)
}
export async function updateFood(foodId, payload) {
    return await axios.put(api + "/" + foodId, payload)
}
export async function ChangeStatus(foodId, status) {
    return await axios.patch(api + "/" + foodId + "/changeStatus/"+status)
}
export async function deleteFood(foodId) {
    return await axios.delete(api + "/" + foodId)
}