import LetterTemplateStore from '@src/pages/templates/store'

function LetterTemplateSettings() {
    const letterTemplateStore = LetterTemplateStore()

    console.log(letterTemplateStore.tergetItem?.content.split(/\{\".*?\"\}/))
    console.log(letterTemplateStore.tergetItem?.content.match(/\{\".*?\"\}/g))
    return (
        <div>LetterTemplate</div>
    )
}

export default LetterTemplateSettings