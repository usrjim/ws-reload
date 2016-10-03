<?php

$config = [
  'origin' => 'http://localhost/',
  'server' => 'ws://localhost:3000/ws',
  'target' => [
    '/path/to/file1',
    '/path/to/file2',
    '/path/to/dir1', 
    '/path/to/dir2', 
  ],
  'log' => '/path/to/log',
];

echo json_encode($config), "\n";

