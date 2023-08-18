async function getData() {
    
    console.log("OK not!")
    const res = await fetch("http://127.0.0.1:8080/");
    const data = await res.json()
    console.log(data)
    
    return data;

}

export default async function HomePage() {
    console.log("dasd")
    const data = await getData()
    return (
        <div>
            Hello World! {data}
        </div>
    )
}