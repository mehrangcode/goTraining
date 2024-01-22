import { create } from 'zustand'
import * as apis from "./api"

export interface TableType {
    id?: string
    name: string
    capacity: number
    status: number
    photos: string[]
}
interface TableStoreType {
    loading: boolean
    list: TableType[]
    targetItem: TableType
    selectTable: (table: TableType) => void
    fetchList: () => void
    create: (payload: TableType) => void
    update: (tableId: string, payload: TableType) => void
    delete: (tableId: string) => void
}

const TableStore = create<TableStoreType>()(
    (set, get) => ({
        loading: false,
        list: [],
        targetItem: undefined,
        fetchList: async () => {
            const res = await apis.getAll()
            set({ list: res.data })
        },
        create: async (payload: TableType) => {
            set({ loading: true })
            try {
                await apis.createTable(payload)
                get().fetchList()
            } catch (error) {
                set({ loading: false })
                throw new Error(error)
            }
        },
        update: async (tableId: string, payload: TableType) => {
            set({ loading: true })
            try {
                const res = await apis.updateTable(tableId, payload)
                const updatedList: TableType[] = JSON.parse(JSON.stringify(get().list))
                const i = updatedList.findIndex(x => x.id === tableId)
                updatedList[i] = res.data
                set({ loading: false, targetItem: undefined, list: updatedList })
            } catch (error) {
                set({ loading: false })
                throw new Error(error)
            }
        },
        selectTable: (table) => {
            set({ targetItem: table })
        },
        delete: async (tableId) => {
            set({ loading: true })
            try {
                await apis.deleteTable(tableId)
                const updatedList = get().list.filter(x => x.id !== tableId)
                set({ loading: false, targetItem: undefined, list: updatedList })
            } catch (error) {
                set({ loading: false })
                throw new Error(error)
            }

        }
    }),
)

export default TableStore