import LetterTemplateStore, { LetterTemplateTargetItemType } from "@src/pages/templates/store"

const defaultItem: LetterTemplateTargetItemType = {
  title: "",
  content: "",
  settings: null
}
function LetterTemplatesForm() {
  const letterTemplateStore = LetterTemplateStore()
  return (
    <div>
      <form>
        <textarea id="content" value={letterTemplateStore.tergetItem?.content} onChange={(e) => {
          e.preventDefault()
          letterTemplateStore.selectItem({
            ...(letterTemplateStore.tergetItem || JSON.parse(JSON.stringify(defaultItem))),
            content: e.target.value
          })
        }} />
      </form>
    </div>
  )
}

export default LetterTemplatesForm