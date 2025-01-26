<script>
    import Sidebar from "../Sidebar.svelte";
    import MainNavbar from "../main-navbar.svelte";
	import '@fontsource/lato';
</script>

<div class="flex flex-col w-screen h-screen">
    <MainNavbar />
    <div class="flex flex-row w-full flex-grow">
        <Sidebar />
        <div class="flex flex-row p-5"></div>
    </div>
</div>