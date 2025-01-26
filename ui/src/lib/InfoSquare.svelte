<script>
    export let color;
    export let iconName;
    export let number;
    export let text;
</script>

<div class="h-0 pb-[100%] rounded-md relative" style="background-color: {color}">
    <div class="flex flex-col p-6 absolute w-full h-full">
        <div class="w-full flex-grow flex justify-end">
            <img src="{iconName}.svg" alt="" class="w-8 h-8">
        </div>
        <p class="font-bold text-5xl text-white">{number}</p>
        <p class="opacity-65 text-lg text-white">{text}</p>
    </div>
</div>