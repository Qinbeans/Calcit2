<script lang="ts">
    import { Parse } from "$lib/ts/wailsjs/go/main/App";
    import { unescapeHTML } from "$lib/ts/helper";

    let innerWidth = 0
    let innerHeight = 0
    let results: { [key: string]: string } | undefined
    let input_value: string = ""

    let function_modal = false
    
    const read_exec = (e: Event) => {
        if (e.target != null) {
            let value = (e.target as Element).innerHTML
            input_value = unescapeHTML(value)
        }
    }
    const submit = (e: Event) => {
        e.preventDefault()
        Parse(input_value).then((res) => {
            //just spits out a number
            let history_eq = document.getElementById("history-eq") as HTMLUListElement
            let li = document.createElement("li")
            li.setAttribute("class", "bg-black/75 pl-[1%] w-full rounded-l overflow-x-clip whitespace-nowrap hover:cursor-pointer")
            li.addEventListener("dblclick", read_exec)
            li.innerHTML = input_value
            history_eq.insertBefore(li, history_eq.firstChild)
            let history_res = document.getElementById("history-res") as HTMLUListElement
            li = document.createElement("li")
            li.setAttribute("class", "bg-black/75 pr-[1%] w-full rounded-r overflow-x-clip whitespace-nowrap hover:cursor-pointer")
            li.innerHTML = "= "+res
            history_res.insertBefore(li, history_res.firstChild)
            if (results == undefined) {
                results = {}
                results[input_value] = res
            } else if (!Object.keys(results).includes(input_value)) {
                results[input_value] = res
            }
        }).catch((err) => {
            alert(err)
            console.log(err)
        })
    }
    const clear = (e: Event) => {
        e.preventDefault()
        let history_eq = document.getElementById("history-eq") as HTMLUListElement
        let history_res = document.getElementById("history-res") as HTMLUListElement
        history_eq.innerHTML = ""
        history_res.innerHTML = ""
        results = undefined
    }
    const save = (e: Event) => {
        e.preventDefault()
        if (results == undefined) {
            return
        }
        let json = JSON.stringify(results)
        let blob = new Blob([json], { type: "application/json" })
        let url = URL.createObjectURL(blob)
        let a = document.createElement("a")
        a.href = url
        a.download = "results.json"
        a.click()
    }
    const load = (e: Event) => {
        e.preventDefault()
        let input = document.createElement("input")
        input.type = "file"
        input.accept = "application/json"
        input.onchange = (e) => {
            if (e.target != null) {
                let file = (e.target as HTMLInputElement).files?.item(0)
                if (file != null) {
                    let reader = new FileReader()
                    reader.onload = (e) => {
                        if (e.target != null) {
                            let json = (e.target as FileReader).result as string
                            let obj = JSON.parse(json)
                            results = obj
                            let history_eq = document.getElementById("history-eq") as HTMLUListElement
                            let history_res = document.getElementById("history-res") as HTMLUListElement
                            history_eq.innerHTML = ""
                            history_res.innerHTML = ""
                            for (let key in results) {
                                let li = document.createElement("li")
                                li.setAttribute("class", "bg-black/75 pl-[1%] w-full rounded-l overflow-x-clip whitespace-nowrap hover:cursor-pointer")
                                li.addEventListener("dblclick", read_exec)
                                li.innerHTML = key
                                history_eq.insertBefore(li, history_eq.firstChild)
                                li = document.createElement("li")
                                li.setAttribute("class", "bg-black/75 pr-[1%] w-full rounded-r overflow-x-clip whitespace-nowrap hover:cursor-pointer")
                                li.innerHTML = "= "+results[key]
                                history_res.insertBefore(li, history_res.firstChild)
                            }
                        }
                    }
                    reader.readAsText(file)
                }
            }
        }
        input.click()
    }
</script>
<svelte:window bind:innerWidth bind:innerHeight />
{#if function_modal}
    <div class="fixed h-screen w-screen z-10 bg-black/40 grid place-content-center">
        <div class="w-full flex flex-row-reverse translate-y-[125%] translate-x-[-0.5%]">
            <button on:click={() => function_modal = !function_modal}>
                <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="red" class="sm:w-[3dvw] sm:h-[3dvw] w-[5dvw] h-[5dvw] hover:fill-red-600">
                    <path stroke-linecap="round" stroke-linejoin="round" d="M9.75 9.75l4.5 4.5m0-4.5l-4.5 4.5M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
                </svg>
            </button>
        </div>
        <div class="h-[75dvh] w-[75dvw] bg-gray-600 rounded p-2 grid grid-cols-3 grid-rows-6">
        </div>
    </div>
{/if}
<div class="flex flex-col gap-1 sm:w-[3dvw] w-[6dvw] fixed left-1 top-1 text-[1dvw]">
    <button class="rounded bg-black/75 hover:bg-black/25 p-1" on:click={save}>
        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="sm:w-[2vw] sm:h-[2vw] w-[4vw] h-[4vw]">
            <path stroke-linecap="round" stroke-linejoin="round" d="M3 16.5v2.25A2.25 2.25 0 005.25 21h13.5A2.25 2.25 0 0021 18.75V16.5M16.5 12L12 16.5m0 0L7.5 12m4.5 4.5V3" />
        </svg>
        Save
    </button>
    <button class="rounded bg-black/75 hover:bg-black/25 p-1" on:click={load}>
        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="sm:w-[2vw] sm:h-[2vw] w-[4vw] h-[4vw]">
            <path stroke-linecap="round" stroke-linejoin="round" d="M3 16.5v2.25A2.25 2.25 0 005.25 21h13.5A2.25 2.25 0 0021 18.75V16.5m-13.5-9L12 3m0 0l4.5 4.5M12 3v13.5" />
          </svg>
          
        Load
    </button>
    <button class="rounded bg-black/75 hover:bg-black/25 p-1 text-red-700" on:click={clear}>
        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="red" class="sm:w-[2vw] sm:h-[2vw] w-[4vw] h-[4vw]">
            <path stroke-linecap="round" stroke-linejoin="round" d="M9.75 9.75l4.5 4.5m0-4.5l-4.5 4.5M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
        </svg>
        Clear
    </button>
</div>
<div class="lg:text-[1.5dvw] sm:text-[3dvw] text-[5dvw]">
    <form id="calc" class="flex gap-1 fixed top-[6.5dvh] left-1/2 transform -translate-x-1/2">
        <input
            type="text"
            placeholder="Enter an equation"
            bind:value={input_value}
            class="rounded overflow-x-clip sm:w-[70vw] w-[65vw] px-2 py-1 bg-black/40 border border-transparent hover:border-green-400"
        >
        <input
            type="submit"
            value="{innerWidth > 640 ? 'Enter' : ">"}"
            class="rounded w-[10vw] px-2 py-1 bg-black/75 hover:bg-black/25 border border-transparent hover:border-green-400"
            on:click={submit}
        >
    </form>
    <div class="top-[12dvh] fixed left-1/2 transform -translate-x-1/2 sm:w-[80dvw] w-[76dvw] flex flex-row-reverse">
        <button class="rounded w-[10vw] px-2 py-1 bg-black/75 hover:bg-black/25 border border-transparent hover:border-green-400" on:click={() => function_modal = !function_modal}>
            {innerWidth > 640 ? "Functions" : "f(x)"}
        </button>
    </div>
    <ul class="flex flex-row fixed left-1/2 transform -translate-x-1/2 top-[20dvh] overflow-y-auto">
        <li>
            <ul class="rounded-l bg-stone-900/75 py-1 pl-[1.5vw] overflow-x-clip flex flex-col-reverse gap-1 sm:w-[75vw] w-[72vw] text-center h-[75vh]" id="history-eq">
            </ul>
        </li>
        <li>
            <ul class="rounded-r bg-stone-900/75 py-1 pr-[1.5vw] overflow-x-clip flex flex-col-reverse gap-1 sm:w-[15vw] w-[12vw] text-center h-[75vh]" id="history-res">
            </ul>
        </li>
    </ul>
</div>
