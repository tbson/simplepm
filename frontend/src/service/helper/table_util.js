class TableUitl {
    static optionToFilter(options) {
        return options.map((option) => ({
            value: option.value,
            text: option.label
        }));
    }

    static optionToTransfer(options) {
        return options.map((option) => ({
            key: `${option.value}`,
            title: option.label,
            description: option.description,
        }));
    }
}

export default TableUitl;
