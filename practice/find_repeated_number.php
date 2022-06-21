<?php
function findRepeatedNumber($numbers) {
    $sameNumbers = [];
    $checked_numbers = array_count_values($numbers);
    foreach ($checked_numbers as $number => $count) {
        if ($count > 1) {
            array_push($sameNumbers, $number);
        }
    }
    return $sameNumbers;

}

print_r(findRepeatedNumber([3,4,2,8,1,2,8]));

?>