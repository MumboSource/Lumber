<script>
    import testIcon from '../../assets/test-img.png'
    import backArrow from '../../assets/left-arrow.svg'
    import { navigate } from 'svelte-routing';
    import Navbar from '../Navbar.svelte';
</script>

<Navbar title="Goat executable" icon={testIcon} backUrl="/exe-list"></Navbar>
<div class="flex flex-col">
    <div class="flex flex-col px-5 py-2">
        <div class="opacity-50">Location</div>
        <div class="text-xl">C:\Users\9281239\Downloads\goat-executable\run.exe</div>
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
<div class="flex p-3 h-16 w-full gap-3 items-center absolute bottom-0" style="border-top: 2px solid #E1E1E1;">
    <button class="w-full h-full transition-all bg-[#4394EB] text-white hover:bg-[#4394ebc4]" on:click={() => navigate("/exe-info")} >
        Set as Start Point
    </button>
</div>