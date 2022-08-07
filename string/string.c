#include <stdio.h>
#include <stdlib.h>

typedef struct string {
  char *data;
  int len;
} string;

string *string_initialize() {
  string *s = malloc(sizeof(string));
  s->data = NULL;
  s->len = 0;
  return s;
}

void string_destroy(string *s) {
  if (s->data) {
    free(s->data);
  }
  free(s);
}

char *string_copy(string *dest, const char *src) {
  for (int i = 0;; i++) {
    dest->data[i] = src[i];
    if (!src[i]) {
      break;
    }
  }
}

size_t string_length(const char *s) {
  size_t len = 0;
  for (;; len++) {
    if (!s[len]) {
      break;
    }
  }
  return len;
}

void string_set(string *s, char *str) {
  int len = string_length(str);
  s->data = malloc(len);
  s->len = len;
  string_copy(s->data, str);
}

int main() {
  string *s = string_initialize();
  string_set(s, "Hello, World!");
  puts(s->data);
  string_destroy(s);
}
