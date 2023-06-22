import { useState } from "react"
import { PostTransaction } from "../wailsjs/go/main/App";

function NewTransaction() {
    const [title, setTitle] = useState<string>('');
    const [amount, setAmount] = useState<number>(0);
    const [categories, setCategories] = useState<string[]>(["compras", "saúde"]);

    const handleSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault()

        let transaction = {
            title: title,
            amount: amount,
            categories: categories.map(e => { return e.trim() }),
        }

        PostTransaction(transaction)
    };

    return (
        <div>
            <h2>New Transaction</h2>
            <form onSubmit={handleSubmit}>
                <label>
                    Title
                    <input type="text" value={title} onChange={e => setTitle(e.target.value)} />
                </label>
                <label>
                    Amount
                    <input type="number" value={amount} onChange={e => setAmount(parseFloat(e.target.value))} />
                </label>
                <label>
                    Categories
                    <input type="text" value={categories} onChange={e => setCategories(e.target.value.split(","))} />
                </label>
                <br />
                <button type="submit">Submit</button>
            </form>
        </div>
    )
}

export default NewTransaction
