import { ConstValues } from "@src/utils/const";
import axios from "axios";

const api = ConstValues.backUri + "/subjects"
export async function getAll() {
    return await axios.get(api)
}

export async function createSubject(payload) {
    return await axios.post(api, payload)
}
export async function updateSubject(subjectId, payload) {
    return await axios.put(api + "/" + subjectId, payload)
}

export async function deleteSubject(subjectId) {
    return await axios.delete(api + "/" + subjectId)
}