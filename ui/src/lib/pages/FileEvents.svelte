<script>
    import Sidebar from "../Sidebar.svelte";
    import Navbar from "../Navbar.svelte";

    let fileEvents = [
        "Test Data",
        "Test Data",
        "Test Data",
        "Test Data",
        "Test Data",
        "Test Data",
        "Test Data"
    ];
    fileEvents = [...fileEvents, ...fileEvents, ...fileEvents, ...fileEvents, ...fileEvents]
</script>

<div class="flex flex-col w-screen h-screen">
    <Navbar title="Goat executable" maximizable={true} />
    <div class="flex flex-row w-full grow">
        <Sidebar />

        <div class="flex flex-col max-h-[calc(100vh-4rem)] w-[calc(100vw-18rem)] overflow-y-scroll">
            <table class="w-full h-full">
                <thead class="bg-[#F2F2F2]">
                    <tr class="border-[#B4B4B4] border-b-2">
                        <th class="py-2 pl-2 w-9"><img src="/question.svg" alt=""></th>
                        <th class="text-[#797979] font-bold border-[#B4B4B4] border-l-2 text-left px-2">TIMESTAMP</th>
                        <th class="text-[#797979] font-bold border-[#B4B4B4] border-l-2 text-left px-2">EVENT TYPE</th>
                        <th class="text-[#797979] font-bold border-[#B4B4B4] border-l-2 text-left px-2">EVENT TARGET</th>
                    </tr>
                </thead>
                <tbody class="font-mono">
                    {#each fileEvents as fileEvent}
                        <tr class="border-[#B4B4B4] border-b-2 hover:bg-[#E5F2FF] hover:text-[#4394EB]">
                            <td class="py-2 pl-2"><img src="/add.svg" alt=""></td>
                            <td class="pl-2">2025-01-26 04:19:27.055 EST</td>
                            <td class="pl-2">File Creation</td>
                            <td class="pl-2">C:\Program Files (x86)\McAfee\mcf.exe</td>
                        </tr>
                    {/each}
                </tbody>
            </table>
        </div>
    </div>
</div>