<?php

// Argon works with PHP 7.2+
$argonPass = "test";
$argonHash = '$argon2i$v=19$m=65536,t=1,p=4$5O+sgJ6jgMUlVVjfFirA4TQ/$NakRk3fpBi0mo/VmUenN+NdTtv6CAeBWzw2KdCqOdgr7';
if(password_verify($argonPass, $argonHash)) {
    echo "Match\n";
} else {
    echo "Nope\n";
}


// Bcrypt works with PHP
$bcryptPass = "test";
$bcryptHash = '$2a$11$UUGBwkNBnqzeAmR2GVFdCeLraf3pYODMBZ.1p7oRGWQMZ97Cfjzny';

if(password_verify($bcryptPass, $bcryptHash)) {
    echo "Match\n";
} else {
    echo "Nope\n";
}
