export interface ModalInterFace {
    open: boolean
    title: string
    onOk: () => void
    onCancel: () => void
    children: any
}
function Modal({
    open,
    title,
    onOk,
    onCancel,
    children
}: ModalInterFace) {
    return (
        <div className="modal">
            <div className="modalHeader">{title}</div>
            <div className="modalBody"></div>
            <div className="modalFooter">
                <button className="btn-primary" onClick={onOk}>Ok</button>
                <button className="btn-primary" onClick={onCancel}>Close</button>
            </div>
        </div>
    )
}

export default Modal