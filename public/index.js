const dog=document.querySelector(".dog-selector")
const jsonBlock=document.querySelector(".json-block")
const fetchBtn=document.querySelector(".fetch-btn")
const clearBtn=document.querySelector(".clear-btn")

async function fetchDog(url,jsonBlock) {
    try {
        const res=await fetch(url)
        const data=await res.json()
        jsonBlock.innerHTML=JSON.stringify(data)
    } catch (error) {
        jsonBlock.innerHTML=error
    }
}


