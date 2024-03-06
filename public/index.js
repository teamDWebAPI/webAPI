const fetchBtn=document.querySelector(".fetch-btn")
const clearBtn=document.querySelector(".clear-btn")
const jsonBlock=document.getElementById("json-block")


async function getDogList(query,keyword, jsonBlock) {
    const serverUrl = "http://localhost:8000/api";
    try {
        const response = await fetch(`${serverUrl}/${query}?keyword=${keyword}`,{ mode: "cors" });
        if (!response.ok) {
            throw new Error(`Error fetching data: ${response.statusText}`);
        }
        const data = await response.json();
        jsonBlock.innerHTML = JSON.stringify(data, null, "\t");
    } catch (error) {
        console.error("Error:", error);
    }
}



fetchBtn.addEventListener("click",(e)=>{
    e.preventDefault()
    removeJson(jsonBlock)
    jsonBlock.classList.add("json-block")
    const dogValue=document.querySelector(".dog-selector").value
    const queryValue=document.querySelector(".query-selector").value
    getDogList(queryValue,dogValue,jsonBlock)
})


clearBtn.addEventListener("click",(e)=>{
    e.preventDefault()
    removeJson(jsonBlock)
})


const removeJson=(jsonBlock)=>{
    jsonBlock.classList.remove("json-block")
    jsonBlock.innerHTML=""
}
