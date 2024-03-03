const dog=document.querySelector(".dog-selector")
const fetchBtn=document.querySelector(".fetch-btn")
const clearBtn=document.querySelector(".clear-btn")
const jsonBlock=document.querySelector(".json-block")

async function getDogList(keyword, jsonBlock) {
    const serverUrl = "http://localhost:8000/api/list";
    try {
        const response = await fetch(`${serverUrl}?keyword=${keyword}`,{ mode: "cors" });
        if (!response.ok) {
            throw new Error(`Error fetching data: ${response.statusText}`);
        }
        const data = await response.json();
        jsonBlock.innerHTML = JSON.stringify(data, null, 2);
    } catch (error) {
        console.error("Error:", error);
    }
}



fetchBtn.addEventListener("click",(e)=>{
    e.preventDefault()
    if(!dog.value){
        alert("please select a dog")
    }

    getDogList(dog.value,jsonBlock)
})


clearBtn.addEventListener("click",(e)=>{
    e.preventDefault()
    jsonBlock.innerHTML=""
})
