package leetcode

import (
	"os"
	"strings"
	"testing"
)

func TestAliasProbe(t *testing.T) {
	if os.Getenv("LEETCODE_LIVE") == "" {
		t.Skip("set LEETCODE_LIVE=1")
	}
	pairs := `binary_tree_inorder_triversal=binary-tree-inorder-traversal
check_every_row_column_contains_all_numbers=check-if-every-row-and-column-contains-all-numbers
contains_duplicate_2=contains-duplicate-ii
count_num_pairs_with_absolute_diff=count-number-of-pairs-with-absolute-difference-k
count_number_of_consistent_strings=count-the-number-of-consistent-strings
count_vowel_substrings=count-vowel-substrings-of-a-string
first_unique_character_in_string=first-unique-character-in-a-string
implement_strstr=find-the-index-of-the-first-occurrence-in-a-string
largest_common_prefix=longest-common-prefix
lowest_common_ancestor_of_bst=lowest-common-ancestor-of-a-binary-search-tree
maximum_units_on_truck=maximum-units-on-a-truck
merge_sorted_list=merge-two-sorted-lists
number_after_a_double_reversal=a-number-after-a-double-reversal
pascals_triangle_2=pascals-triangle-ii
pivot_index=find-pivot-index
roman_to_interger=roman-to-integer
running_sum=running-sum-of-1d-array
target_Indices_after_sorting_array=find-target-indices-after-sorting-array
x_matrix=check-if-matrix-is-x-matrix
add_two_numbers_ll=add-two-numbers
check_if_there_valid_partition_for_the_array=check-if-there-is-a-valid-partition-for-the-array
construct_b_tree_pre_inorder_traversal=construct-binary-tree-from-preorder-and-inorder-traversal
jump_game_4=jump-game-iv
linked_list_cycle_2=linked-list-cycle-ii
longest_increasing_subsequence=longest-increasing-subsequence
max_area_cake_piece=maximum-area-of-a-piece-of-cake-after-horizontal-and-vertical-cuts
max_area_of_island=max-area-of-island
merge_intervals=merge-intervals
three_sum_closest=3sum-closest
triangular_sum_of_an_array=find-triangular-sum-of-an-array`

	c := NewClient()
	for _, line := range strings.Split(pairs, "\n") {
		parts := strings.SplitN(strings.TrimSpace(line), "=", 2)
		m, found, err := c.FetchQuestionMeta(parts[1])
		if err != nil || !found {
			t.Errorf("MISS %-46s -> %-60s err=%v", parts[0], parts[1], err)
			continue
		}
		t.Logf("OK   %-46s -> %-4d %-8s %s", parts[0], m.Number, m.Difficulty, m.Slug)
	}
}
