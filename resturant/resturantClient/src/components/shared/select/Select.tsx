import { useMemo, useRef, useState } from "react"

type SelectProps = {
    options: { [key: string]: string }[]
    value: string
    onChange: (value: string) => void
    disabled?: boolean
    valueKey?: string
    displayKey?: string
}
function Select({
    options,
    valueKey = "id",
    displayKey = 'label',
    value,
    onChange
}: SelectProps) {
    const [showBox, setShowBox] = useState(false)
    const [searchValue, setSearchValue] = useState("")
    const inputRef = useRef(null)
    const _value = useMemo(() => {
        let result = ""
        if (options.find(x => x[valueKey] === value)) {
            return options.find(x => x[valueKey] === value)[displayKey]
        }
        return result
    }, [value])
    return (
        <div className="select"
            tabIndex={0}
            onClick={() => {setShowBox(true)}}
            onFocus={() => {
                setShowBox(true)
                if (inputRef.current) {
                    inputRef.current.focus()
                }
            }}
            onBlur={() => setShowBox(false)}
        >
            <span className="valueSpan" style={{
                opacity: searchValue?.length > 0 ? .3 : 1,
                width: searchValue?.length > 0 ? 0 : "max-content",
                overflowX: searchValue?.length > 0 ? "hidden" : "visible",
            }}>{_value}</span>
            <input
                ref={inputRef}
                type="text" className="selectInput"
                onFocus={(e) => e.target.select()}
                value={searchValue} onChange={(e) => {
                    setShowBox(true)
                    setSearchValue(e.target.value)
                }} />
            <div className={showBox ? "selectBox open" : "selectBox"}>
                {options.filter(x => !searchValue || x[displayKey].toLowerCase().includes(searchValue.toLowerCase())).map(item => {
                    return <div
                        className="option"
                        key={item[valueKey]}
                        onClick={(e) => {
                            e.stopPropagation()
                            setSearchValue("")
                            if (onChange) {
                                onChange(item[valueKey])
                                setShowBox(false)
                            }
                        }}>{item[displayKey]}</div>
                })}
            </div>
        </div>
    )
}

export default Select