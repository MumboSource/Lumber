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
</script>

<div class="flex flex-col w-screen h-screen">
    <Navbar title="McAfee Pro Crack" maximizable={true} />
    <div class="flex flex-row w-full grow">
        <Sidebar />
        <div class="flex flex-col grow max-h-[calc(100vh-4rem)] overflow-y-scroll">
            <div class="flex w-full h-8 items-center bg-[#F2F2F2] border-b-2 border-[#B4B4B4]">
                <div class="w-8 h-full flex border-r-2 border-[#B4B4B4] px-2 items-center justify-center">?</div>
                <div class="grow-[6] h-full border-r-2 border-[#B4B4B4] px-2 flex items-center">TIMESTAMP</div>
                <div class="grow-[2] h-full border-r-2 border-[#B4B4B4] px-2 flex items-center">EVENT TYPE</div>
                <div class="grow-[8] h-full border-r-2 border-[#B4B4B4] px-2 flex items-center">EVENT TARGET</div>
            </div>
            {#each fileEvents as fileEvent}
                <div class="flex w-full h-8 items-center border-b-2 border-[#B4B4B4]">
                    <div class="w-8 h-full flex px-2 items-center justify-center">?</div>
                    <div class="grow-[6] h-full px-2 flex items-center ">{new Date()}</div>
                    <div class="grow-[2] h-full px-2 flex items-center">EVENT TYPE</div>
                    <div class="grow-[8] h-full px-2 flex items-center">EVENT TARGET</div>
                </div>
            {/each}
        </div>
    </div>
</div>