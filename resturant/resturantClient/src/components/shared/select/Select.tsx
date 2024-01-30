import { ReactElement, useMemo, useRef, useState } from "react"

type SelectProps = {
    options: any[]
    value: string | string[]
    onChange: (value: string) => void
    disabled?: boolean
    valueKey?: string
    displayKey?: string
    mode?: "single" | "multiple"
}
function Select({
    options,
    valueKey = "id",
    displayKey = 'label',
    mode = "single",
    value,
    onChange
}: SelectProps) {
    const [showBox, setShowBox] = useState(false)
    const [searchValue, setSearchValue] = useState("")
    const inputRef = useRef(null)
    const _value = useMemo(() => {
        let result = value
        if (value && mode === "single" && options.find(x => x[valueKey] === value)) {
            return options.find(x => x[valueKey] === value)[displayKey]
        }
        if (value && mode === "multiple") {
            const _values: ReactElement[] = [];
            (value as string[]).forEach(v => {
                const t = options.find(x => x[valueKey] === v)
                if (t) {
                    _values.push(<span className="multipleOption">{t[displayKey]}</span>)
                } else {
                    _values.push(<span className="multipleOption">{v}</span>)
                }
            })
            return _values
        }
        return result
    }, [value, mode])
    return (
        <div className="select"
            tabIndex={0}
            onClick={() => { setShowBox(true) }}
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
                overflow: "hidden",
                display: "flex",
                gap: "0.25rem"
            }}>{
                    typeof _value == "undefined" ? null : _value
                }</span>
            <input
                style={{ width: ((searchValue?.length || 1) * 8) + "px" }}
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