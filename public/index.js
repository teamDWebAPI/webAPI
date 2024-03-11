const fetchBtn=document.querySelector(".fetch-btn")
const clearBtn=document.querySelector(".clear-btn")
const jsonBlock=document.getElementById("json-block")


async function getDogList(breed,jsonBlock) {
    const serverUrl = "http://localhost:8000/api";
    try {
        console.log(`${serverUrl}/images?breed=${breed}`)
        const response = await fetch(`${serverUrl}/images?breed=${breed}`,{ mode: "cors" });
        if (!response.ok) {
            throw new Error(`Error fetching data: ${response.statusText}`);
        }
        const data = await response.json();
        const img_element = document.createElement('img');
        img_element.src = data["message"][0];
        jsonBlock.appendChild(img_element);
    } catch (error) {
        console.error("Error:", error);
    }
}



fetchBtn.addEventListener("click",(e)=>{
    e.preventDefault()
    removeJson(jsonBlock)
    jsonBlock.classList.add("json-block")
    const dogValue=document.querySelector(".dog-selector").value
    // const subbreedValue=document.querySelector(".subbreed-selector").value
    // const countValue=document.querySelector(".count").value
    getDogList(dogValue,jsonBlock)
})


clearBtn.addEventListener("click",(e)=>{
    e.preventDefault()
    removeJson(jsonBlock)
})


const removeJson=(jsonBlock)=>{
    jsonBlock.classList.remove("json-block")
    jsonBlock.innerHTML=""
}
