import { create } from 'zustand'
import * as UserApis from "./api"

interface UserType {
    name: string
    email: string
    password: string
}
interface UserStoreType {
    loading: boolean
    list: UserType[]
    fetchList: () => void
    create: (payload: UserType) => void
}

const UserStore = create<UserStoreType>()(
    (set, get) => ({
        loading: false,
        list: [],
        fetchList: async () => {
            const res = await UserApis.getAll()
            set({ list: res.data })
        },
        create: async (payload: UserType) => {
            set({ loading: true })
            try {
                await UserApis.create(payload)
                get().fetchList()
            } catch (error) {
                throw new Error(error as any)
            }
        }
    }),
)

export default UserStore