import { create } from "zustand";


export interface LetterTemplateTargetItemType {
    title: string
    content: string
    settings: {
        options: {id: string; label: string; type: string}
    }
}
type LetterTemplateStoreType = {
    loading: boolean
    tergetItem: LetterTemplateTargetItemType
    selectItem: (item: LetterTemplateTargetItemType) => void
}
const LetterTemplateStore = create<LetterTemplateStoreType>((set) => ({
    loading: false,
    tergetItem: undefined,
    selectItem: (item: LetterTemplateTargetItemType) => {
        set({tergetItem: item})
    }
}))

export default LetterTemplateStore