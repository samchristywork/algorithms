#include <stdio.h>

#define SWAP(x, y) \
  int t = x;       \
  x = y;           \
  y = t;

int compare(int a, int b) {
  return a < b;
}

int abs_compare(int a, int b) {
  if (a < 0)
    a *= -1;
  if (b < 0)
    b *= -1;
  return a < b;
}

void quicksort(int *arr, int len, int (*compare)(int, int)) {
  if (len < 2) {
    return;
  }

  int l = 0;
  int r = len - 1;
  while (1) {
    int m = arr[len / 2];
    while (compare(arr[l], m)) {
      l++;
    }
    while (compare(m, arr[r])) {
      r--;
    }
    if (l >= r) {
      break;
    }

    SWAP(arr[l], arr[r]);

    l++;
    r--;
  }

  quicksort(arr, l, compare);
  quicksort(arr + l, len - l, compare);
}

void print_array(int *arr, int len) {
  for (int i = 0; i < len; i++) {
    printf("%d ", arr[i]);
  }
  printf("\n");
}

int main() {
  int arr[] = {
      27, 71, 5, 8, 44, 38, -4, 33, 21, 18, 77, 33, -33, 20, 62, 5, 48, 77, 20, -83, 36, 18};
  int len = sizeof(arr) / sizeof(int);

  print_array(arr, len);

  quicksort(arr, len, compare);
  print_array(arr, len);

  quicksort(arr, len, abs_compare);
  print_array(arr, len);
}
