import LetterTemplatesForm from '@src/components/templates/letters/LetterTemplatesForm'
import LetterTemplateSettings from '@src/components/templates/letters/Settings'

function LetterTemplatesPage() {
    return (
        <div className='letterTemplate'>
            <section className="view">
                <div className="templates">
                    templates
                </div>
                <LetterTemplatesForm />
                <LetterTemplateSettings />
            </section>
        </div>
    )
}

export default LetterTemplatesPage