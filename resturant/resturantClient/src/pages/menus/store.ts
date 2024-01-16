import { create } from 'zustand'
import * as apis from "./api"

export interface sectionFoodType {
    price: number
    food_id: string
}
export interface SectionType {
    id?: string
    title: string
    description: string
    foods: sectionFoodType[]
}
export interface MenuType {
    id?: string
    title: string
    description: string
    status: number
    sections: SectionType[]
}
export interface MenuStoreType {
    loading: boolean
    list: MenuType[]
    targetItem: MenuType
    selectMenu: (menu: MenuType) => void
    fetchList: () => void
    create: (payload: MenuType) => void
    update: (menuId: string, payload: MenuType) => void
    delete: (menuId: string) => void
}

const MenuStore = create<MenuStoreType>()(
    (set, get) => ({
        loading: false,
        list: [],
        targetItem: undefined,
        fetchList: async () => {
            const res = await apis.getAll()
            set({ list: res.data })
        },
        create: async (payload: MenuType) => {
            set({ loading: true })
            try {
                await apis.createMenu(payload)
                get().fetchList()
            } catch (error) {
                set({ loading: false })
                throw new Error(error)
            }
        },
        update: async (menuId: string, payload: MenuType) => {
            set({ loading: true })
            try {
                const res = await apis.updateMenu(menuId, payload)
                const updatedList: MenuType[] = JSON.parse(JSON.stringify(get().list))
                const i = updatedList.findIndex(x => x.id === menuId)
                updatedList[i] = res.data
                set({ loading: false, targetItem: undefined, list: updatedList })
            } catch (error) {
                set({ loading: false })
                throw new Error(error)
            }
        },
        selectMenu: (menu) => {
            set({ targetItem: menu })
        },
        delete: async (menuId) => {
            set({ loading: true })
            try {
                await apis.deleteMenu(menuId)
                const updatedList = get().list.filter(x => x.id !== menuId)
                set({ loading: false, targetItem: undefined, list: updatedList })
            } catch (error) {
                set({ loading: false })
                throw new Error(error)
            }

        }
    }),
)

export default MenuStore