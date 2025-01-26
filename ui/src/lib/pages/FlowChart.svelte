<script>
    import Navbar from "../Navbar.svelte";
    import Sidebar from "../Sidebar.svelte";
    import { Node, Svelvet, Minimap, Controls } from 'svelvet';
</script>

<div class="flex flex-col w-screen h-screen">
    <Navbar title="Goat executable" maximizable={true} />
    <div class="flex flex-row w-full grow">
        <Sidebar />
        <div class="p-3 grow h-full flex items-center justify-center">
            <Svelvet width={1100-288-24} height={700-64-24}>
                <Node />
                <Node id="alpha" bgColor="red" label="Default Node" />
            </Svelvet>
        </div>
    </div>
</div>