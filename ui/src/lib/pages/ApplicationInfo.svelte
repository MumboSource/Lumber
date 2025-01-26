<script>
	import InfoSquare from './../InfoSquare.svelte';
    import Sidebar from "../Sidebar.svelte";
    import MainNavbar from "../Navbar.svelte";
	import '@fontsource/lato';
</script>

<div class="flex flex-col w-screen h-screen">
    <MainNavbar />
    <div class="flex flex-row w-full">
        <Sidebar />
        <div class="flex flex-row p-5 gap-6 flex-grow overflow-y-scroll max-h-[calc(100vh-4rem)]">
            <div class="flex flex-col gap-6 flex-1">
                <p class="font-mono bg-[#F0F0F0] py-3 px-4 rounded-md text-3xl w-full">McAfee_CRACK.exe</p>
                {#each {length: 15} as i}
                    <div>
                        <p class="opacity-50">Date Created</p>
                        <p class="text-xl">12/13/24 10:20:30 EST</p>
                    </div>
                {/each}
            </div>
            <div class="flex flex-col gap-6 flex-1">
                <div class="grid grid-cols-2 grid-rows-2 gap-2">
                    <InfoSquare color="#4394EB" number="12" text="Files Accessed" iconName="accessed"/>
                    <InfoSquare color="#1875D9" number="2" text="Executions" iconName="execution"/>
                    <InfoSquare color="#125CAC" number="17" text="Actions Incurred" iconName="actions"/>
                    <InfoSquare color="#074281" number="4" text="Child Processes" iconName="processes"/>
                </div>
                {#each {length: 10} as i}
                    <div>
                        <p class="opacity-50">Date Created</p>
                        <p class="text-xl">12/13/24 10:20:30 EST</p>
                    </div>
                {/each}
            </div>
        </div>
    </div>
</div>