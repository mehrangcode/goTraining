import { create } from 'zustand'
import * as apis from "./api"

export interface FoodCategoryType {
    id?: string
    title: string
    description: string
    status: number
    photos: string
}
interface FoodCategoryStoreType {
    loading: boolean
    list: FoodCategoryType[]
    targetItem: FoodCategoryType
    selectFood: (food: FoodCategoryType) => void
    fetchList: () => void
    create: (payload: FoodCategoryType) => void
    update: (foodId: string, payload: FoodCategoryType) => void
    changeStatus: (foodId: string, status: number) => void
    delete: (foodId: string) => void
}

const FoodCategoryStore = create<FoodCategoryStoreType>()(
    (set, get) => ({
        loading: false,
        list: [],
        targetItem: undefined,
        fetchList: async () => {
            const res = await apis.getAll()
            set({ list: res.data })
        },
        create: async (payload: FoodCategoryType) => {
            set({ loading: true })
            try {
                await apis.createFood(payload)
                get().fetchList()
            } catch (error) {
                set({ loading: false })
                throw new Error(error)
            }
        },
        update: async (foodId: string, payload: FoodCategoryType) => {
            set({ loading: true })
            try {
                const res = await apis.updateFood(foodId, payload)
                const updatedList: FoodCategoryType[] = JSON.parse(JSON.stringify(get().list))
                const i = updatedList.findIndex(x => x.id === foodId)
                updatedList[i] = res.data
                set({ loading: false, targetItem: undefined, list: updatedList })
            } catch (error) {
                set({ loading: false })
                throw new Error(error)
            }
        },
        selectFood: (food) => {
            set({ targetItem: food })
        },
        changeStatus: async (foodId: string, status: number) => {
            set({ loading: true })
            try {
                await apis.ChangeStatus(foodId, status)
                const updatedList: FoodCategoryType[] = JSON.parse(JSON.stringify(get().list))
                const i = updatedList.findIndex(x => x.id === foodId)
                updatedList[i].status = status
                set({ loading: false, targetItem: undefined, list: updatedList })
            } catch (error) {
                set({ loading: false })
                throw new Error(error)
            }
        },
        delete: async (foodId) => {
            set({ loading: true })
            try {
                await apis.deleteFood(foodId)
                const updatedList = get().list.filter(x => x.id !== foodId)
                set({ loading: false, targetItem: undefined, list: updatedList })
            } catch (error) {
                set({ loading: false })
                throw new Error(error)
            }

        }
    }),
)

export default FoodCategoryStore