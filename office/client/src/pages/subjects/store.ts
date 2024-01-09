import { create } from 'zustand'
import * as apis from "./api"

interface SubjectType {
    id?: string
    label: string
    archive?: boolean
}
interface SubjectStoreType {
    loading: boolean
    list: SubjectType[]
    targetItem: SubjectType
    selectSubject: (subject: SubjectType) => void
    fetchList: () => void
    create: (payload: SubjectType) => void
    update: (subjectId: string, payload: SubjectType) => void
    delete: (subjectId: string) => void
}

const SubjectStore = create<SubjectStoreType>()(
    (set, get) => ({
        loading: false,
        list: [],
        targetItem: undefined,
        fetchList: async () => {
            const res = await apis.getAll()
            set({ list: res.data })
        },
        create: async (payload: SubjectType) => {
            set({ loading: true })
            try {
                await apis.createSubject(payload)
                get().fetchList()
            } catch (error) {
                set({ loading: false })
                throw new Error(error)
            }
        },
        update: async (subjectId: string, payload: SubjectType) => {
            set({ loading: true })
            try {
                const res = await apis.updateSubject(subjectId, payload)
                const updatedList: SubjectType[] = JSON.parse(JSON.stringify(get().list))
                const i = updatedList.findIndex(x => x.id === subjectId)
                updatedList[i] = res.data
                set({ loading: false, targetItem: undefined, list: updatedList })
            } catch (error) {
                set({ loading: false })
                throw new Error(error)
            }
        },
        selectSubject: (subject) => {
            set({ targetItem: subject })
        },
        delete: async (subjectId) => {
            set({ loading: true })
            try {
                await apis.deleteSubject(subjectId)
                const updatedList = get().list.filter(x => x.id !== subjectId)
                set({ loading: false, targetItem: undefined, list: updatedList })
            } catch (error) {
                set({ loading: false })
                throw new Error(error)
            }

        }
    }),
)

export default SubjectStore