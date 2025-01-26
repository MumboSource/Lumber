<script>
    import Sidebar from "../Sidebar.svelte";
    import Navbar from "../Navbar.svelte";
    import ExecutableDetails from "./ExecutableDetails.svelte";

    let executables = [
        {icon: "/chrome.png", name: "Google Chrome", path: "C:\\Program Files\\Google\\Chrome\\chrome.exe"},
        {icon: "/chrome.png", name: "Google Chrome", path: "C:\\Program Files\\Google\\Chrome\\chrome.exe"},
        {icon: "/chrome.png", name: "Google Chrome", path: "C:\\Program Files\\Google\\Chrome\\chrome.exe"},
        {icon: "/chrome.png", name: "Google Chrome", path: "C:\\Program Files\\Google\\Chrome\\chrome.exe"},
        {icon: "/chrome.png", name: "Google Chrome", path: "C:\\Program Files\\Google\\Chrome\\chrome.exe"},
        {icon: "/chrome.png", name: "Google Chrome", path: "C:\\Program Files\\Google\\Chrome\\chrome.exe"},
        {icon: "/chrome.png", name: "Google Chrome", path: "C:\\Program Files\\Google\\Chrome\\chrome.exe"},
        {icon: "/chrome.png", name: "Google Chrome", path: "C:\\Program Files\\Google\\Chrome\\chrome.exe"},
        {icon: "/chrome.png", name: "Google Chrome", path: "C:\\Program Files\\Google\\Chrome\\chrome.exe"},
        {icon: "/chrome.png", name: "Google Chrome", path: "C:\\Program Files\\Google\\Chrome\\chrome.exe"},
        {icon: "/chrome.png", name: "Google Chrome", path: "C:\\Program Files\\Google\\Chrome\\chrome.exe"},
        {icon: "/chrome.png", name: "Google Chrome", path: "C:\\Program Files\\Google\\Chrome\\chrome.exe"},
        {icon: "/chrome.png", name: "Google Chrome", path: "C:\\Program Files\\Google\\Chrome\\chrome.exe"},
        {icon: "/chrome.png", name: "Google Chrome", path: "C:\\Program Files\\Google\\Chrome\\chrome.exe"},
        {icon: "/chrome.png", name: "Google Chrome", path: "C:\\Program Files\\Google\\Chrome\\chrome.exe"},
        {icon: "/chrome.png", name: "Google Chrome", path: "C:\\Program Files\\Google\\Chrome\\chrome.exe"},
        {icon: "/chrome.png", name: "Google Chrome", path: "C:\\Program Files\\Google\\Chrome\\chrome.exe"},
        {icon: "/chrome.png", name: "Google Chrome", path: "C:\\Program Files\\Google\\Chrome\\chrome.exe"},
        {icon: "/chrome.png", name: "Google Chrome", path: "C:\\Program Files\\Google\\Chrome\\chrome.exe"},
        {icon: "/chrome.png", name: "Google Chrome", path: "C:\\Program Files\\Google\\Chrome\\chrome.exe"},
        {icon: "/chrome.png", name: "Google Chrome", path: "C:\\Program Files\\Google\\Chrome\\chrome.exe"},
        {icon: "/chrome.png", name: "Google Chrome", path: "C:\\Program Files\\Google\\Chrome\\chrome.exe"},
        {icon: "/chrome.png", name: "Google Chrome", path: "C:\\Program Files\\Google\\Chrome\\chrome.exe"},
        {icon: "/chrome.png", name: "Google Chrome", path: "C:\\Program Files\\Google\\Chrome\\chrome.exe"},
        {icon: "/chrome.png", name: "Google Chrome", path: "C:\\Program Files\\Google\\Chrome\\chrome.exe"},
        {icon: "/chrome.png", name: "Google Chrome", path: "C:\\Program Files\\Google\\Chrome\\chrome.exe"},
        {icon: "/chrome.png", name: "Google Chrome", path: "C:\\Program Files\\Google\\Chrome\\chrome.exe"},
        {icon: "/chrome.png", name: "Google Chrome", path: "C:\\Program Files\\Google\\Chrome\\chrome.exe"},
        {icon: "/chrome.png", name: "Google Chrome", path: "C:\\Program Files\\Google\\Chrome\\chrome.exe"},
        {icon: "/chrome.png", name: "Google Chrome", path: "C:\\Program Files\\Google\\Chrome\\chrome.exe"},
        {icon: "/chrome.png", name: "Google Chrome", path: "C:\\Program Files\\Google\\Chrome\\chrome.exe"},
    ]

    let currentExecutable = executables[0];

    let modalOpen = false;
</script>

{#if modalOpen}
<div class="fixed inset-0 bg-black bg-opacity-50">
    <div class="fixed top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 bg-white rounded-sm">
        <Navbar title={currentExecutable.name} icon={currentExecutable.icon} modalMode on:click={() => modalOpen = false}></Navbar>
        <div class="flex flex-col">
            <div class="flex flex-col px-5 py-2">
                <div class="opacity-50">Location</div>
                <div class="text-xl">C:\Users\9281239\Downloads\mcafee-pro\run.exe</div>
            </div>
            <div class="flex flex-col px-5 py-2">
                <div class="opacity-50">Size</div>
                <div class="text-xl">1.3 YB</div>
            </div>
            <div class="flex flex-col px-5 py-2">
                <div class="opacity-50">Date Created</div>
                <div class="text-xl">August 32, 1738 25:00:23 FM EST</div>
            </div>
        </div>
    </div>
</div>
{/if}

<div class="flex flex-col w-screen h-screen">
    <Navbar title="McAfee Pro Crack" maximizable={true} />
    <div class="flex flex-row w-full grow">
        <Sidebar />
        <div class="flex flex-col max-h-[calc(100vh-4rem)] w-[calc(100vw-18rem)] overflow-y-scroll">
            <table class="w-full h-full">
                <thead class="bg-[#F2F2F2]">
                    <tr class="border-[#B4B4B4] border-b-2">
                        <th class="py-2 pl-2 w-9"><img src="/question.svg" alt=""></th>
                        <th class="text-[#797979] font-bold border-[#B4B4B4] border-l-2 text-left px-2">PROGRAM NAME</th>
                        <th class="text-[#797979] font-bold border-[#B4B4B4] border-l-2 text-left px-2">EXECUTABLE PATH</th>
                    </tr>
                </thead>
                <tbody>
                    {#each executables as exec}
                        <tr class="border-[#E0E0E0] border-b-2 hover:bg-[#E5F2FF] hover:text-[#4394EB] cursor-pointer" on:click={() => {currentExecutable = exec; modalOpen = true}}>
                            <td class="py-2 pl-2"><img src={exec.icon} alt="" class="w-5 h-5"></td>
                            <td class="pl-2">{exec.name}</td>
                            <td class="pl-2 flex flex-row justify-between">
                                <p>{exec.path}</p>
                                <img src="/chevron-right.svg" alt="" class="hover:backdrop-brightness-90 rounded-md"/>
                            </td>
                        </tr>
                    {/each}
                </tbody>
            </table>
        </div>
    </div>
</div>