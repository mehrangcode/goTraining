import { create } from 'zustand'
import * as apis from "./api"

interface UserType {
    id?: string
    name: string
    email: string
    password: string
}
interface UserStoreType {
    loading: boolean
    list: UserType[]
    targetItem: UserType
    selectUser: (user: UserType) => void
    fetchList: () => void
    create: (payload: UserType) => void
    update: (userId: string, payload: UserType) => void
    delete: (userId: string) => void
}

const UserStore = create<UserStoreType>()(
    (set, get) => ({
        loading: false,
        list: [],
        targetItem: undefined,
        fetchList: async () => {
            const res = await apis.getAll()
            set({ list: res.data })
        },
        create: async (payload: UserType) => {
            set({ loading: true })
            try {
                await apis.createUser(payload)
                get().fetchList()
            } catch (error) {
                set({ loading: false })
                throw new Error(error)
            }
        },
        update: async (userId: string, payload: UserType) => {
            set({ loading: true })
            try {
                const res = await apis.updateUser(userId, payload)
                const updatedList: UserType[] = JSON.parse(JSON.stringify(get().list))
                const i = updatedList.findIndex(x => x.id === userId)
                updatedList[i] = res.data
                set({ loading: false, targetItem: undefined, list: updatedList })
            } catch (error) {
                set({ loading: false })
                throw new Error(error)
            }
        },
        selectUser: (user) => {
            set({ targetItem: user })
        },
        delete: async (userId) => {
            set({ loading: true })
            try {
                await apis.deleteUser(userId)
                const updatedList = get().list.filter(x => x.id !== userId)
                set({ loading: false, targetItem: undefined, list: updatedList })
            } catch (error) {
                set({ loading: false })
                throw new Error(error)
            }

        }
    }),
)

export default UserStore