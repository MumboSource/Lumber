<script>
    import Navbar from "../Navbar.svelte";
    import Sidebar from "../Sidebar.svelte";
    import { navigate } from "svelte-routing";

    import rightArrow from '../../assets/right-arrow.svg'

    let users = [
        "Ruadhan KellerMcLeer",
        "Ruadhan KellerMcLeer",
        "Ruadhan KellerMcLeer",
        "Ruadhan KellerMcLeer",
        "Ruadhan KellerMcLeer",
        "Tim Timbers",
    ];
</script>

<div class="flex flex-col w-screen h-screen">
    <Navbar title="Goat executable" maximizable={true} />
    <div class="flex flex-row w-full grow">
        <Sidebar />
        <div class="flex flex-col w-full">
            <div class="flex p-6 gap-6">
                <div class="flex flex-col w-3/5 h-full max-h-96">
                    <div class="flex justify-between border-2 border-[#E0E0E0] p-3 items-center h-16">
                        <div class="text-2xl font-bold">Users</div>
                        <div>Open Settings</div>
                    </div>
                    <div class="overflow-y-scroll no-scrollbar">
                        {#each users as user}
                            <div class="flex border-x-2 border-b-2 border-[#E0E0E0] p-3 items-center h-16">
                                <img alt="pfp" />
                                <div class="flex flex-col grow">
                                    <div class=" text-xl">{user}</div>
                                    <div class="text-sm opacity-50">Last Logon: 11/12/24</div>
                                </div>
                                <button on:click={() => console.log("RUN SOME POWERSHELL")} >
                                    <img src={rightArrow} alt="select" class="nextButton w-12 rounded-[5px] transition-all" />
                                </button>
                            </div>
                        {/each}
                    </div>
                </div>
                <div class="flex flex-col w-2/5 max-h-96 justify-between">
                    <div class="border-2 border-[#E0E0E0] w-full h-28 flex items-center p-3">
                        <div class="flex flex-col h-full bg-slate-100 justify-center p-3 rounded-md">
                            <div class="text-xl font-bold">16YB <span class="font-normal text-sm">Memory</span></div>
                            <div class="italic opacity-50">22% in use</div>
                        </div>
                    </div>
                    <div class="border-2 border-[#E0E0E0] w-full h-28 flex items-center p-3">
                        <div class="flex flex-col h-full bg-slate-100 justify-center p-3 rounded-md">
                            <div class="text-xl font-bold">16YB <span class="font-normal text-sm">Memory</span></div>
                            <div class="italic opacity-50">22% in use</div>
                        </div>
                    </div>
                    <div class="border-2 border-[#E0E0E0] w-full h-28 flex items-center p-3">
                        <div class="flex flex-col h-full bg-slate-100 justify-center p-3 rounded-sm">
                            <div class="text-xl font-bold">16YB <span class="font-normal text-sm">Memory</span></div>
                            <div class="italic opacity-50">22% in use</div>
                        </div>
                    </div>
                </div>
            </div>
            <div class="px-6">
                <div class="flex flex-col w-full">
                    <div class="flex grow">
                        <div class="flex flex-col">
                            <div class="opacity-50">Device Name</div>
                            <div class="text-xl">Sigma Laptop</div>
                        </div>
                    </div>
                    <div class="flex grow">
                        <div class="flex flex-col">
                            <div class="opacity-50">Device Name</div>
                            <div class="text-xl">Sigma Laptop</div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</div>

<style>
    .nextButton:hover {
        background-color: #d9d9d9;
    }
</style>